package render

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// sitemapURL — один <url> елемент sitemap.xml (schema sitemaps.org/schemas/sitemap/0.9).
type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

// sitemapURLSet — кореневий елемент <urlset> sitemap.xml.
type sitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

// collectSitemapURLs обходить дерево груп (той самий порядок обходу, що й
// Generate()) і збирає канонічні URL усіх сторінок сайту — головна, країни,
// міста, заклади — тими самими хелперами *URL, що й canonicalOf/навігація
// (задача 3.5), обгорнутими в absURL: sitemap.xml вимагає абсолютних адрес.
func collectSitemapURLs(groups []domain.CountryPlaceGroup) []sitemapURL {
	urls := []sitemapURL{{Loc: absURL(IndexURL())}}

	for _, country := range groups {
		urls = append(urls, sitemapURL{Loc: absURL(CountryURL(country.Country))})

		for _, city := range country.Cities {
			urls = append(urls, sitemapURL{Loc: absURL(CityURL(city.City))})

			for _, place := range city.Places {
				u := sitemapURL{Loc: absURL(PlaceURL(place))}
				// AddedAt заповнений для кожного заклада (CLAUDE.md, Status) —
				// використовуємо як <lastmod>; нульове значення (якби
				// траплялось) пропускаємо, щоб не писати фіктивну дату 0001-01-01.
				if !place.AddedAt.IsZero() {
					u.LastMod = place.AddedAt.Format("2006-01-02")
				}
				urls = append(urls, u)
			}
		}
	}

	return urls
}

// writeSitemap будує sitemap.xml (schema sitemaps.org) за деревом груп і
// записує його в корінь outputDir (задача 3.5). Помилка — через error, як і
// решта пайплайну Generate, без паніки.
func writeSitemap(groups []domain.CountryPlaceGroup, outputDir string) error {
	set := sitemapURLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  collectSitemapURLs(groups),
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := enc.Encode(set); err != nil {
		return fmt.Errorf("render: encode sitemap: %w", err)
	}
	buf.WriteByte('\n')

	path := filepath.Join(outputDir, "sitemap.xml")
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("render: write %s: %w", path, err)
	}

	return nil
}
