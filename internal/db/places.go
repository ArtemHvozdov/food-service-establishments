package db

import (
	"fmt"
	"sort"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// Places завантажує сирі дані (1.2), мапить кожен запис у domain.Place (1.3)
// і повертає їх згрупованими по країнах і містах (internal_docs/task_01.md, 1.4).
// Помилки завантаження/мапінгу — це збій даних на етапі білда, тому панікуємо
// (як mustDate), а не повертаємо error: сигнатура Places() лишається старою.
func Places() ([]domain.CountryPlaceGroup, error) {
	dtos, err := loadPlaceDTOs()
	if err != nil {
		return nil, fmt.Errorf("db: завантаження даних: %w", err)
	}
	places := make([]domain.Place, 0, len(dtos))
	for _, dto := range dtos {
		place, err := dtoToPlace(dto)
		if err != nil {
			return nil, fmt.Errorf("db: мапінг заведення: %w", err)
		}
		places = append(places, place)
	}
	return groupPlaces(places), nil
}

// cityKey — унікальний ключ міста в межах усього набору: alias міста сам по
// собі не унікальний між країнами (напр. однаковий alias у двох країнах).
type cityKey struct {
	countryAlias string
	cityAlias    string
}

// groupPlaces групує плаский список заведень у []domain.CountryPlaceGroup:
// країна -> місто -> заклади. У межах міста дедуплікує alias і сортує заклади
// за назвою; країни й міста сортуються за alias — усе для відтворюваності
// порядку між білдами.
func groupPlaces(places []domain.Place) []domain.CountryPlaceGroup {
	countries := map[string]domain.Country{}
	cities := map[cityKey]domain.City{}
	byCity := map[cityKey][]domain.Place{}
	cityAliasesByCountry := map[string]map[string]bool{}

	for _, p := range places {
		key := cityKey{countryAlias: p.City.Country.Alias, cityAlias: p.City.Alias}

		countries[key.countryAlias] = p.City.Country
		cities[key] = p.City
		byCity[key] = append(byCity[key], p)

		if cityAliasesByCountry[key.countryAlias] == nil {
			cityAliasesByCountry[key.countryAlias] = map[string]bool{}
		}
		cityAliasesByCountry[key.countryAlias][key.cityAlias] = true
	}

	countryAliases := make([]string, 0, len(countries))
	for alias := range countries {
		countryAliases = append(countryAliases, alias)
	}
	sort.Strings(countryAliases)

	groups := make([]domain.CountryPlaceGroup, 0, len(countryAliases))
	for _, countryAlias := range countryAliases {
		cityAliases := make([]string, 0, len(cityAliasesByCountry[countryAlias]))
		for alias := range cityAliasesByCountry[countryAlias] {
			cityAliases = append(cityAliases, alias)
		}
		sort.Strings(cityAliases)

		cityGroups := make([]domain.CityPlaceGroup, 0, len(cityAliases))
		for _, cityAlias := range cityAliases {
			key := cityKey{countryAlias: countryAlias, cityAlias: cityAlias}
			cityPlaces := byCity[key]

			sort.SliceStable(cityPlaces, func(i, j int) bool {
				if cityPlaces[i].Name != cityPlaces[j].Name {
					return cityPlaces[i].Name < cityPlaces[j].Name
				}
				if cityPlaces[i].Alias != cityPlaces[j].Alias {
					return cityPlaces[i].Alias < cityPlaces[j].Alias
				}
				// Фінальний тай-брейк по GooglePlaceID (непустий і унікальний
				// для всіх заведень): порядок при повному збігу Name/Alias не
				// має залежати від позиції запису у вхідному файлі, інакше
				// призначення суфікса -2 "пливе" між білдами (internal_docs/task_01.md, 1.4.1).
				return cityPlaces[i].GooglePlaceID < cityPlaces[j].GooglePlaceID
			})
			dedupeAliases(cityPlaces)

			cityGroups = append(cityGroups, domain.CityPlaceGroup{
				City:   cities[key],
				Places: cityPlaces,
			})
		}

		groups = append(groups, domain.CountryPlaceGroup{
			Country: countries[countryAlias],
			Cities:  cityGroups,
		})
	}

	return groups
}

// dedupeAliases додає детермінований суфікс -2, -3, ... до повторюваних
// Place.Alias у межах уже відсортованого зрізу заведень одного міста.
// `used` враховує і вихідні, і вже згенеровані alias, тому кандидат base-N
// перевіряється на зайнятість перед призначенням — інакше можна випадково
// зіткнутися з "натуральним" alias виду base-2, що вже є серед заведень
// (internal_docs/task_01.md, 1.4.1).
func dedupeAliases(places []domain.Place) {
	used := make(map[string]bool, len(places))
	for _, p := range places {
		used[p.Alias] = true
	}

	seen := map[string]int{}
	for i := range places {
		base := places[i].Alias
		seen[base]++
		if seen[base] == 1 {
			continue // перше входження лишається без суфікса
		}

		n := seen[base]
		candidate := fmt.Sprintf("%s-%d", base, n)
		for used[candidate] {
			n++
			candidate = fmt.Sprintf("%s-%d", base, n)
		}

		places[i].Alias = candidate
		used[candidate] = true
		seen[base] = n
	}
}
