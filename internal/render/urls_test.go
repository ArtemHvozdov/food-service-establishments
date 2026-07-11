package render

import (
	"strings"
	"testing"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

func testPlace() domain.Place {
	return domain.Place{
		Alias: "borshch-krakow",
		Name:  "Борщ",
		City: domain.City{
			Alias: "krakow",
			Name:  "Краків",
			Country: domain.Country{
				Alias: "poland",
				Name:  "Польща",
			},
		},
	}
}

// TestURLsAreRootRelative перевіряє, що всі URL, повернуті хелперами,
// починаються з "/" (internal_docs/task_01.md, 1.6 DoD): відносні посилання
// заборонені через копію public/[country]/index.html на іншій глибині.
func TestURLsAreRootRelative(t *testing.T) {
	place := testPlace()
	city := place.City
	country := city.Country

	urls := map[string]string{
		"index":   IndexURL(),
		"country": CountryURL(country),
		"city":    CityURL(city),
		"place":   PlaceURL(place),
	}

	for name, url := range urls {
		if !strings.HasPrefix(url, "/") {
			t.Errorf("%s: URL %q не починається з /", name, url)
		}
	}
}

// TestPathsAndURLsAreConsistent перевіряє, що шлях файлу і URL узгоджені між
// собою: URL — це шлях без префікса "public".
func TestPathsAndURLsAreConsistent(t *testing.T) {
	place := testPlace()
	city := place.City
	country := city.Country

	if got, want := IndexFilePath(), "public/index.html"; got != want {
		t.Errorf("IndexFilePath() = %q, want %q", got, want)
	}
	if got, want := IndexURL(), "/"; got != want {
		t.Errorf("IndexURL() = %q, want %q", got, want)
	}

	flatFile, indexFile := CountryFilePaths(country)
	if got, want := flatFile, "public/poland.html"; got != want {
		t.Errorf("CountryFilePaths() flatFile = %q, want %q", got, want)
	}
	if got, want := indexFile, "public/poland/index.html"; got != want {
		t.Errorf("CountryFilePaths() indexFile = %q, want %q", got, want)
	}
	if got, want := CountryURL(country), "/poland.html"; got != want {
		t.Errorf("CountryURL() = %q, want %q", got, want)
	}
	if "public/"+strings.TrimPrefix(CountryURL(country), "/") != flatFile {
		t.Errorf("CountryURL() %q неузгоджений з плоским файлом %q", CountryURL(country), flatFile)
	}

	if got, want := CityFilePath(city), "public/poland/krakow.html"; got != want {
		t.Errorf("CityFilePath() = %q, want %q", got, want)
	}
	if got, want := CityURL(city), "/poland/krakow.html"; got != want {
		t.Errorf("CityURL() = %q, want %q", got, want)
	}
	if "public/"+strings.TrimPrefix(CityURL(city), "/") != CityFilePath(city) {
		t.Errorf("CityURL() неузгоджений з CityFilePath()")
	}

	if got, want := PlaceFilePath(place), "public/poland/krakow/borshch-krakow.html"; got != want {
		t.Errorf("PlaceFilePath() = %q, want %q", got, want)
	}
	if got, want := PlaceURL(place), "/poland/krakow/borshch-krakow.html"; got != want {
		t.Errorf("PlaceURL() = %q, want %q", got, want)
	}
	if "public/"+strings.TrimPrefix(PlaceURL(place), "/") != PlaceFilePath(place) {
		t.Errorf("PlaceURL() неузгоджений з PlaceFilePath()")
	}
}
