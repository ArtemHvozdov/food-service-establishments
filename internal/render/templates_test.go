package render

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// placeWithAllLinks — заведення з усіма трьома соцпосиланнями + сайтом, щоб
// перевірити, що жодне з них не пропускається шаблоном.
func placeWithAllLinks() domain.Place {
	p := testPlace()
	p.URL = "https://borshch.example"
	p.FacebookURL = "https://facebook.com/borshch"
	p.InstagramURL = "https://instagram.com/borshch"
	return p
}

// placeWithOnlyURL — заведення лише з сайтом, щоб перевірити, що відсутні
// Facebook/Instagram не рендерять порожні <a>.
func placeWithOnlyURL() domain.Place {
	p := testPlace()
	p.URL = "https://borshch.example"
	return p
}

func testCityGroup(places ...domain.Place) domain.CityPlaceGroup {
	return domain.CityPlaceGroup{City: places[0].City, Places: places}
}

func testCountryGroup(places ...domain.Place) domain.CountryPlaceGroup {
	return domain.CountryPlaceGroup{
		Country: places[0].City.Country,
		Cities:  []domain.CityPlaceGroup{testCityGroup(places...)},
	}
}

// TestTemplatesCompileAndExecute перевіряє, що всі чотири шаблони компілюються
// (package-level template.Must у init вже б впав інакше) і виконуються без
// помилок на валідних доменних даних (internal_docs/task_01.md, 1.7 DoD).
func TestTemplatesCompileAndExecute(t *testing.T) {
	place := placeWithAllLinks()
	city := testCityGroup(place)
	country := testCountryGroup(place)
	index := []domain.CountryPlaceGroup{country}

	tests := []struct {
		name string
		data any
		exec func(data any) (*bytes.Buffer, error)
	}{
		{"place", place, func(data any) (*bytes.Buffer, error) {
			var buf bytes.Buffer
			err := placeTemplate.Execute(&buf, data)
			return &buf, err
		}},
		{"city", city, func(data any) (*bytes.Buffer, error) {
			var buf bytes.Buffer
			err := cityTemplate.Execute(&buf, data)
			return &buf, err
		}},
		{"country", country, func(data any) (*bytes.Buffer, error) {
			var buf bytes.Buffer
			err := countryTemplate.Execute(&buf, data)
			return &buf, err
		}},
		{"index", index, func(data any) (*bytes.Buffer, error) {
			var buf bytes.Buffer
			err := indexTemplate.Execute(&buf, data)
			return &buf, err
		}},
	}

	h1Re := regexp.MustCompile(`(?i)<h1[ >]`)
	titleRe := regexp.MustCompile(`(?i)<title>([^<]+)</title>`)
	canonicalRe := regexp.MustCompile(`(?i)<link rel="canonical" href="([^"]+)">`)

	for _, tc := range tests {
		buf, err := tc.exec(tc.data)
		if err != nil {
			t.Fatalf("%s: execute failed: %v", tc.name, err)
		}
		out := buf.String()

		if !strings.Contains(out, `<html lang="uk">`) {
			t.Errorf("%s: no <html lang=\"uk\">", tc.name)
		}
		if got := h1Re.FindAllString(out, -1); len(got) != 1 {
			t.Errorf("%s: expected exactly one <h1>, got %d", tc.name, len(got))
		}
		m := titleRe.FindStringSubmatch(out)
		if m == nil || strings.TrimSpace(m[1]) == "" {
			t.Errorf("%s: missing or empty <title>", tc.name)
		}
		cm := canonicalRe.FindStringSubmatch(out)
		if cm == nil || !strings.HasPrefix(cm[1], "/") {
			t.Errorf("%s: missing or non-root canonical link", tc.name)
		}
	}
}

// TestPlaceTemplateOmitsMissingLinks перевіряє, що відсутня соцмережа не дає
// порожнього <a>: сторінка з лише сайтом не повинна містити посилань на
// Facebook/Instagram.
func TestPlaceTemplateOmitsMissingLinks(t *testing.T) {
	var buf bytes.Buffer
	if err := placeTemplate.Execute(&buf, placeWithOnlyURL()); err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	out := buf.String()

	if !strings.Contains(out, "https://borshch.example") {
		t.Errorf("expected site link in output, got: %s", out)
	}
	if strings.Contains(out, "Facebook") {
		t.Errorf("did not expect Facebook link in output, got: %s", out)
	}
	if strings.Contains(out, "Instagram") {
		t.Errorf("did not expect Instagram link in output, got: %s", out)
	}
}

// TestPlaceTemplateIncludesAllLinks перевіряє, що всі три соцпосилання
// (сайт/Facebook/Instagram) рендеряться, коли вони присутні.
func TestPlaceTemplateIncludesAllLinks(t *testing.T) {
	var buf bytes.Buffer
	place := placeWithAllLinks()
	if err := placeTemplate.Execute(&buf, place); err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	out := buf.String()

	for _, want := range []string{place.URL, place.FacebookURL, place.InstagramURL} {
		if !strings.Contains(out, want) {
			t.Errorf("expected output to contain %q, got: %s", want, out)
		}
	}
}

// TestCityTemplateLinksToPlaces перевіряє, що список закладів міста лінкує на
// сторінку кожного заведення через href з хелпера 1.6.
func TestCityTemplateLinksToPlaces(t *testing.T) {
	place := placeWithAllLinks()
	city := testCityGroup(place)

	var buf bytes.Buffer
	if err := cityTemplate.Execute(&buf, city); err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	out := buf.String()

	want := `href="` + PlaceURL(place) + `"`
	if !strings.Contains(out, want) {
		t.Errorf("expected place link %q in output, got: %s", want, out)
	}
}
