package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// geoBBox — грубий bounding box країни (мін/макс широта і довгота), ключ —
// Country.Alias. Використовується лише щоб зловити заклад, що "відлетів" в
// іншу країну (напр. переплутані координати скрапера) — не точна перевірка
// адреси.
type geoBBox struct {
	minLat, maxLat float64
	minLng, maxLng float64
}

var countryBBox = map[string]geoBBox{
	"bulgaria": {41.2, 44.3, 22.3, 28.7},
	"cyprus":   {34.5, 35.8, 32.2, 34.7},
	"czechia":  {48.5, 51.1, 12.0, 18.9},
	"germany":  {47.2, 55.1, 5.8, 15.1},
	"spain":    {35.9, 43.9, -9.4, 4.4},
	"hungary":  {45.7, 48.6, 16.1, 22.9},
	"poland":   {49.0, 54.9, 14.1, 24.2},
	"portugal": {36.9, 42.2, -9.6, -6.1},
	"romania":  {43.6, 48.3, 20.2, 29.7},
	"slovakia": {47.7, 49.6, 16.8, 22.6},
	"turkiye":  {35.8, 42.1, 25.6, 44.8},
}

// hasValidURLScheme перевіряє, що непорожній URL починається з http:// або
// https:// (ловить друкарські помилки на кшталт "hhttps://"). Порожній рядок
// вважається допустимим — поле не заповнене.
func hasValidURLScheme(url string) bool {
	return url == "" || strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// Validate обходить дерево країна -> місто -> заклад, повернуте Places(), і
// перевіряє інваріанти, на які спирається рендер (internal_docs/task_01.md,
// 1.5). Це чекпоінт над уже зібраним результатом — не дублює логіку
// мапера/дедупу (1.3, 1.4), а лише підтверджує, що вона не порушена.
// Повертає об'єднану помилку з усіма знайденими порушеннями (nil, якщо їх
// немає).
func Validate(groups []domain.CountryPlaceGroup) error {
	var errs []error

	seenCountryAlias := map[string]bool{}

	for _, country := range groups {
		switch {
		case country.Country.Alias == "":
			errs = append(errs, fmt.Errorf("db: empty Country.Alias (country name %q)", country.Country.Name))
		case seenCountryAlias[country.Country.Alias]:
			errs = append(errs, fmt.Errorf("db: duplicate Country.Alias %q", country.Country.Alias))
		}
		seenCountryAlias[country.Country.Alias] = true

		seenCityAlias := map[string]bool{}

		for _, city := range country.Cities {
			switch {
			case city.City.Alias == "":
				errs = append(errs, fmt.Errorf("db: empty City.Alias (city name %q, country %q)", city.City.Name, country.Country.Alias))
			case seenCityAlias[city.City.Alias]:
				errs = append(errs, fmt.Errorf("db: duplicate City.Alias %q in country %q", city.City.Alias, country.Country.Alias))
			}
			seenCityAlias[city.City.Alias] = true

			seenPlaceAlias := map[string]bool{}

			for _, place := range city.Places {
				if place.Name == "" {
					errs = append(errs, fmt.Errorf("db: empty Place.Name (alias %q, city %q/%q)", place.Alias, country.Country.Alias, city.City.Alias))
				}

				switch {
				case place.Alias == "":
					errs = append(errs, fmt.Errorf("db: empty Place.Alias (name %q, city %q/%q)", place.Name, country.Country.Alias, city.City.Alias))
				case seenPlaceAlias[place.Alias]:
					errs = append(errs, fmt.Errorf("db: duplicate Place.Alias %q in city %q/%q", place.Alias, country.Country.Alias, city.City.Alias))
				}
				seenPlaceAlias[place.Alias] = true

				if place.URL == "" && place.GoogleMapsURL == "" && place.FacebookURL == "" && place.InstagramURL == "" {
					errs = append(errs, fmt.Errorf("db: place %q (alias %q, city %q/%q) has no outgoing link (URL/GoogleMapsURL/FacebookURL/InstagramURL all empty)", place.Name, place.Alias, country.Country.Alias, city.City.Alias))
				}

				bbox, ok := countryBBox[country.Country.Alias]
				if !ok {
					errs = append(errs, fmt.Errorf("db: no bounding box defined for country %q (place %q, alias %q)", country.Country.Alias, place.Name, place.Alias))
				} else {
					lat, lng := place.Location.Latitude, place.Location.Longitude
					if lat < bbox.minLat || lat > bbox.maxLat || lng < bbox.minLng || lng > bbox.maxLng {
						errs = append(errs, fmt.Errorf("db: place %q (alias %q, country %q) has coordinates (%g, %g) outside country bounding box", place.Name, place.Alias, country.Country.Alias, lat, lng))
					}
				}

				urlFields := []struct {
					name string
					url  string
				}{
					{"URL", place.URL},
					{"GoogleMapsURL", place.GoogleMapsURL},
					{"MenuURL", place.MenuURL},
					{"PhotoURL", place.PhotoURL},
					{"FacebookURL", place.FacebookURL},
					{"InstagramURL", place.InstagramURL},
				}
				for _, f := range urlFields {
					if !hasValidURLScheme(f.url) {
						errs = append(errs, fmt.Errorf("db: place %q (alias %q, city %q/%q) has %s with invalid scheme: %q", place.Name, place.Alias, country.Country.Alias, city.City.Alias, f.name, f.url))
					}
				}
			}
		}
	}

	return errors.Join(errs...)
}
