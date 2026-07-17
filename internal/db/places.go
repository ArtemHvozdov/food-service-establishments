package db

import (
	"fmt"
	"sort"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

func Places() []domain.CountryPlaceGroup {
	return []domain.CountryPlaceGroup{
		{
			Country: domain.Bulgaria,
			Cities: []domain.CityPlaceGroup{
				{
					City: domain.BulgariaVarna,
					Places: []domain.Place{
						{
							Alias:           "bg-varna-stefania-restaurant",
							PreviousAliases: nil,
							City:            domain.BulgariaVarna,
							Name:            "Stefania Restaurant",
							Address:         "бул. Княз Борис I 126, 9010 Варна, България",
							Type:            domain.Restaurant,
							URL:             "https://stefania.bg/bg/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/stefania_rest/?hl=bg/",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/Restaurant+Stefania+DE/@52.5350217,13.6081279,17z/data=!3m1!4b1!4m6!3m5!1s0x47a84b00192cc3b5:0xc4b85b5d593e979a!8m2!3d52.5350217!4d13.6081279!16s%2Fg%2F11lmkc2wh9?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJb-6thO5VpEAR6upb0CSxOS4",
							Location: domain.Location{
								Latitude:  52.5350217,
								Longitude: 13.6081279,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.BulgariaSofia,
					Places: []domain.Place{
						{
							Alias:           "bg-sofia-faina-ukraina",
							PreviousAliases: nil,
							City:            domain.BulgariaSofia,
							Name:            "Файна Украина",
							Address:         "ул. Странджа 44, 1303 София, България",
							Type:            domain.Restaurant,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/faina_ukraina_2026/",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/%D0%A4%D0%B0%D0%B9%D0%BD%D0%B0+%D0%AE%D0%BA%D1%80%D0%B0%D0%B9%D0%BD%D0%B0/@42.6991315,23.3118459,17z/data=!3m1!4b1!4m6!3m5!1s0x40aa85bca639d98d:0x2a50d8f4492ee026!8m2!3d42.6991315!4d23.3118459!16s%2Fg%2F11z5641r8_?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJjdk5pryFqkARJuAuSfTYUCo",
							Location: domain.Location{
								Latitude:  42.6991315,
								Longitude: 23.3118459,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
						{
							Alias:           "bg-sofia-bistro-portokal",
							PreviousAliases: nil,
							City:            domain.BulgariaSofia,
							Name:            "Бистро Портокал",
							Address:         "кв. Банишора, ул. Опалченска 53, 1233 София, България",
							Type:            domain.Bistro,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/%D0%91%D0%B8%D1%81%D1%82%D1%80%D0%B8+,,%D0%9F%D0%BE%D1%80%D1%82%D0%BE%D0%BA%D0%B0%D0%BB%E2%80%9D/@42.7105402,23.3145189,17z/data=!4m16!1m9!3m8!1s0x40aa8555f0abbba5:0x8badbd7ce286cb9a!2z0JHQuNGB0YLRgNC4ICws0J_QvtGA0YLQvtC60LDQu-KAnQ!8m2!3d42.7105402!4d23.3145189!9m1!1b1!16s%2Fg%2F11sj5839rz!3m5!1s0x40aa8555f0abbba5:0x8badbd7ce286cb9a!8m2!3d42.7105402!4d23.3145189!16s%2Fg%2F11sj5839rz?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJpbur8FWFqkARmsuG4ny9rYs",
							Location: domain.Location{
								Latitude:  42.7105402,
								Longitude: 23.3145189,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
					},
				},
			},
		},
		{
			Country: domain.Cyprus,
			Cities: []domain.CityPlaceGroup{
				{
					City: domain.CyprusKouklia,
					Places: []domain.Place{
						{
							Alias:           "cy-kouklia-eda-family-cafe",
							PreviousAliases: nil,
							City:            domain.CyprusKouklia,
							Name:            "Eda Family Cafe",
							Address:         "Kouklia 8500, Cyprus",
							Type:            domain.Cafe,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/eda_family_cafe/",
							FacebookURL:     "https://www.facebook.com/people/EDA_family_cafe/100089954059264/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Eda+family+cafe/@34.7088816,32.5747878,17z/data=!3m1!4b1!4m6!3m5!1s0x14e71bc500341d0b:0x74caf78ede5ee4c9!8m2!3d34.7088816!4d32.5747878!16s%2Fg%2F11kjglygq_?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJCx00AMUb5xQRyeRe3o73ynQ",
							Location: domain.Location{
								Latitude:  34.7088816,
								Longitude: 32.5747878,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CyprusLarnaca,
					Places: []domain.Place{
						{
							Alias:           "cy-larnaca-berezka-store-cuisine",
							PreviousAliases: nil,
							City:            domain.CyprusLarnaca,
							Name:            "Berezka store & cuisine",
							Address:         "Lordou Vyrona 35, Larnaka 6023, Cyprus",
							Type:            domain.Cafe,
							URL:             "https://berezka.cy/en",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/berezka.cy/",
							FacebookURL:     "https://www.facebook.com/berezka.cy/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Berezka+store+%26+cuisine/@34.9162087,33.6349798,14z/data=!4m10!1m2!2m1!1sberezka+store+%26+cuisine+larnaca!3m6!1s0x14e083ef0d373bd9:0xb10c41f8a98927ac!8m2!3d34.9162087!4d33.6349798!15sCh9iZXJlemthIHN0b3JlICYgY3Vpc2luZSBsYXJuYWNhWiEiH2JlcmV6a2Egc3RvcmUgJiBjdWlzaW5lIGxhcm5hY2GSAQRkZWxpmgEjQ2haRFNVaE5NRzluUzBWSlEwRm5TVU5vZFZCdVVGaDNFQUXgAQD6AQUIhgEQKg!16s%2Fg%2F11k4q4rqt5?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ2Ts3De-D4BQRrCeJqfhBDLE",
							Location: domain.Location{
								Latitude:  34.9162087,
								Longitude: 33.6349798,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CyprusNicosia,
					Places: []domain.Place{
						{
							Alias:           "cy-nicosia-berezka-store-cuisine",
							PreviousAliases: nil,
							City:            domain.CyprusNicosia,
							Name:            "BEREZKA store & cuisine",
							Address:         "Platonos 1, Strovolos 2040, Nicosia, Cyprus",
							Type:            domain.Cafe,
							URL:             "https://berezka.cy/en",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/berezka.cy/",
							FacebookURL:     "https://www.facebook.com/berezka.cy/",
							GoogleMapsURL:   "https://www.google.com/maps/place/BEREZKA+store+%26+cuisine/@35.1388686,33.3418928,13z/data=!4m10!1m2!2m1!1sberezka+store+%26+cuisine+nicosia!3m6!1s0x14de1754cdfd4927:0x5de04cc2b998e1f8!8m2!3d35.1388686!4d33.3418928!15sCh9iZXJlemthIHN0b3JlICYgY3Vpc2luZSBuaWNvc2lhWiEiH2JlcmV6a2Egc3RvcmUgJiBjdWlzaW5lIG5pY29zaWGSAQ1tZWFsX3Rha2Vhd2F5mgEjQ2haRFNVaE5NRzluUzBWSlEwRm5TVVJFYlU5UU5XWm5FQUXgAQD6AQQIJhBA!16s%2Fg%2F11stfbb1tv?entry=ttu&g_ep=EgoyMDI2MDYyOS4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJJ0n9zVQX3hQR-OGYucJM4F0",
							Location: domain.Location{
								Latitude:  35.1388686,
								Longitude: 33.3418928,
							},
							FoundingYear:  0,
							Description:   "",
							Dishes:        nil,
							BorschPrice:   "",
							LargePortions: false,
							HasDelivery:   false,
							DonatesToZSU:  nil,
							Articles:      nil,
							YoutubeVideos: nil,
							AddedAt:       mustDate("2026-07-02"),
							PromotedUntil: nil,
						},
					},
				},
			},
		},
		{
			Country: domain.Czechia,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Germany,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Spain,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Hungary,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Poland,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Portugal,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Romania,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Slovakia,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
		{
			Country: domain.Turkiye,
			Cities:  []domain.CityPlaceGroup{
				// ...
			},
		},
	}
}

// PlacesOld завантажує сирі дані (1.2), мапить кожен запис у domain.Place (1.3)
// і повертає їх згрупованими по країнах і містах (internal_docs/task_01.md, 1.4).
// Помилки завантаження/мапінгу — це збій даних на етапі білда, тому панікуємо
// (як mustDate), а не повертаємо error: сигнатура PlacesOld() лишається старою.
func PlacesOld() ([]domain.CountryPlaceGroup, error) {
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
