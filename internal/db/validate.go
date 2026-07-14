package db

import (
	"errors"
	"fmt"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// Validate обходить дерево країна -> місто -> заклад, повернуте PlacesOld(), і
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
			}
		}
	}

	return errors.Join(errs...)
}
