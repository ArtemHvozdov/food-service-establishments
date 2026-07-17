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
			Cities: []domain.CityPlaceGroup{
				{
					City: domain.CzechiaBrno,
					Places: []domain.Place{
						{
							Alias:           "cz-brno-defiliada",
							PreviousAliases: nil,
							City:            domain.CzechiaBrno,
							Name:            "Defiliada",
							Address:         "Libušina tř. 2, 623 00 Brno-Kohoutovice, Česká republika",
							Type:            domain.Restaurant,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "https://www.facebook.com/people/Defiliada-Defiliada/61577147900521/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Defiliada/@49.1954462,16.5400605,17z/data=!4m16!1m9!3m8!1s0x4712970053a25df9:0x9e9935778eda47cb!2sDefiliada!8m2!3d49.1954462!4d16.5400605!9m1!1b1!16s%2Fg%2F11xh3dg65b!3m5!1s0x4712970053a25df9:0x9e9935778eda47cb!8m2!3d49.1954462!4d16.5400605!16s%2Fg%2F11xh3dg65b?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ-V2iUwCXEkcRy0fajnc1mZ4",
							Location: domain.Location{
								Latitude:  49.1954462,
								Longitude: 16.5400605,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaJihlava,
					Places: []domain.Place{
						{
							Alias:           "cz-jihlava-u-cimbora",
							PreviousAliases: nil,
							City:            domain.CzechiaJihlava,
							Name:            "Přátelské Bistro U Cimbora",
							Address:         "Březinova 4690/144, 586 01 Jihlava, Česká republika",
							Type:            domain.Bistro,
							URL:             "https://ucimbora.cz/#",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "https://www.facebook.com/people/P%C5%99%C3%A1telsk%C3%A9-Bistro-U-Cimbora/100079996854694/",
							GoogleMapsURL:   "https://www.google.com/maps/place/P%C5%99%C3%A1telsk%C3%A9+Bistro+U+Cimbora/@49.4018175,15.6021138,17z/data=!3m1!4b1!4m6!3m5!1s0x470d1b007af02b3f:0x77347b5e75d098cf!8m2!3d49.4018175!4d15.6021138!16s%2Fg%2F11wwj46y3s?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJPyvwegAbDUcRz5jQdV57NHc",
							Location: domain.Location{
								Latitude:  49.4018175,
								Longitude: 15.6021138,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaKarlovyVary,
					Places: []domain.Place{
						{
							Alias:           "cz-karlovy-vary-ukraina-restaurace",
							PreviousAliases: nil,
							City:            domain.CzechiaKarlovyVary,
							Name:            "Ukrajna restaurace",
							Address:         "Západní 13, 360 01 Karlovy Vary, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://restauraceukrajina.cz/cs/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/Restaurace+%22Ukrajina%22+Karlovy+Vary,+%D0%A0%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD+%22%D0%A3%D0%BA%D1%80%D0%B0%D1%97%D0%BD%D0%B0%22+%D0%9A%D0%B0%D1%80%D0%BB%D0%BE%D0%B2%D1%8B+%D0%92%D0%B0%D1%80%D1%8B,+Restaurant+%22Ukraine%22+in+Karlsbad+-+ukrajinsk%C3%A1+restaurace/@50.2294048,12.8638128,17z/data=!3m1!4b1!4m6!3m5!1s0x47a099f64c1ea48d:0xa6626c3b1506dcb8!8m2!3d50.2294048!4d12.8638128!16s%2Fg%2F11nn3nkz_8?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJjaQeTPaZoEcRuNwGFTtsYqY",
							Location: domain.Location{
								Latitude:  50.2294048,
								Longitude: 12.8638128,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaLiberec,
					Places: []domain.Place{
						{
							Alias:           "cz-liberec-po-nashomu",
							PreviousAliases: nil,
							City:            domain.CzechiaLiberec,
							Name:            "Po-Nashomu",
							Address:         "Vaňurova 458/13, 460 07 Liberec, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://po-nashomu.cz/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/po_nashomu_/",
							FacebookURL:     "https://www.facebook.com/CleopatraLiberec",
							GoogleMapsURL:   "https://www.google.com/maps/place/Po-Nashomu/@50.7642331,15.0472529,17z/data=!3m1!4b1!4m6!3m5!1s0x470937001ca79a9f:0x3d3a2ca51cb59122!8m2!3d50.7642331!4d15.0472529!16s%2Fg%2F11n4tcdf73?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJn5qnHAA3CUcRIpG1HKUsOj0",
							Location: domain.Location{
								Latitude:  50.7642331,
								Longitude: 15.0472529,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaMladaBoleslav,
					Places: []domain.Place{
						{
							Alias:           "cz-mlada-boleslav-clock-cafe",
							PreviousAliases: nil,
							City:            domain.CzechiaMladaBoleslav,
							Name:            "Clock Café",
							Address:         "Českobratrské nám. 57/3, 293 01 Mladá Boleslav, Česká republika",
							Type:            domain.Cafe,
							URL:             "https://clockcafe.cz/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/kavarnaclockcafe/",
							FacebookURL:     "https://www.facebook.com/kavarnaclockcafe/",
							GoogleMapsURL:   "https://www.google.com/maps/place/CLOCK+CAF%C3%89/@50.4140055,14.9048264,17z/data=!3m1!4b1!4m6!3m5!1s0x4709554d3f3e75eb:0x1793494cc33fae3e!8m2!3d50.4140055!4d14.9048264!16s%2Fg%2F11lj13k5pw?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ63U-P01VCUcRPq4_w0xJkxc",
							Location: domain.Location{
								Latitude:  50.4140055,
								Longitude: 14.9048264,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaOlomouc,
					Places: []domain.Place{
						{
							Alias:           "cz-olomouc-simple-dumpl",
							PreviousAliases: nil,
							City:            domain.CzechiaOlomouc,
							Name:            "Simple Dumpl",
							Address:         "Denisova 310/13, 779 00 Olomouc, Česká republika",
							Type:            domain.Bistro,
							URL:             "https://simpledumpl.com/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/simple.dumpl/",
							FacebookURL:     "https://www.facebook.com/p/Simple-Dumpl-100066411355254/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Simple+Dumpl/@49.5955804,17.2544667,17z/data=!4m16!1m9!3m8!1s0x47124fe273b4b8fb:0x4e02c79a37fee123!2sSimple+Dumpl!8m2!3d49.5955804!4d17.2544667!9m1!1b1!16s%2Fg%2F11kq7tc9lj!3m5!1s0x47124fe273b4b8fb:0x4e02c79a37fee123!8m2!3d49.5955804!4d17.2544667!16s%2Fg%2F11kq7tc9lj?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ-7i0c-JPEkcRI-H-N5rHAk4",
							Location: domain.Location{
								Latitude:  49.5955804,
								Longitude: 17.2544667,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaPardubice,
					Places: []domain.Place{
						{
							Alias:           "cz-pardubice-berehynia",
							PreviousAliases: nil,
							City:            domain.CzechiaPardubice,
							Name:            "Berehynia Ukrajinská kuchyně",
							Address:         "Náměstí Republiky 53, 530 02 Pardubice, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://bleriotpardubice.cz/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "https://www.facebook.com/bleriotpardubice/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Berehynia+Ukrajinsk%C3%A1+kuchyn%C4%9B/@50.0241904,15.7596832,17z/data=!4m16!1m9!3m8!1s0x470dcd0c277016f3:0x84d8e59e320c4a37!2sBerehynia+Ukrajinsk%C3%A1+kuchyn%C4%9B!8m2!3d50.0241904!4d15.7596832!9m1!1b1!16s%2Fg%2F11l5s392gk!3m5!1s0x470dcd0c277016f3:0x84d8e59e320c4a37!8m2!3d50.0241904!4d15.7596832!16s%2Fg%2F11l5s392gk?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ8xZwJwzNDUcRN0oMMp7l2IQ",
							Location: domain.Location{
								Latitude:  50.0241904,
								Longitude: 15.7596832,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaPraha,
					Places: []domain.Place{
						{
							Alias:           "cz-praha-bistro-panska",
							PreviousAliases: nil,
							City:            domain.CzechiaPraha,
							Name:            "Bistro Panská (Soloha)",
							Address:         "Panská 1, 110 00 Praha 1-Nové Město, Česká republika",
							Type:            domain.Bistro,
							URL:             "https://bistropanska.choiceqr.com/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/Bistro+Pansk%C3%A1/@50.0850006,14.4265641,17z/data=!3m1!4b1!4m6!3m5!1s0x470b95906896b343:0xea8dc47248e9e0a5!8m2!3d50.0850006!4d14.4265641!16s%2Fg%2F11tjs8tyth?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJQ7OWaJCVC0cRpeDpSHLEjeo",
							Location: domain.Location{
								Latitude:  50.0850006,
								Longitude: 14.4265641,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
						{
							Alias:           "cz-praha-bistro-soloha",
							PreviousAliases: nil,
							City:            domain.CzechiaPraha,
							Name:            "Bistro Soloha",
							Address:         "J. Masaryka 36, 120 00 Praha 2-Vinohrady, Česká republika",
							Type:            domain.Bistro,
							URL:             "https://soloha.uds.app/c",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/soloha_bistro/",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/Bistro+Soloha/@50.07385,14.4314889,3334m/data=!3m1!1e3!4m10!1m2!2m1!1sbistro+soloha+prague!3m6!1s0x470b95acb17d668d:0xc4b603ddcf9f9cca!8m2!3d50.0718449!4d14.4395247!15sChRiaXN0cm8gc29sb2hhIHByYWd1ZVoWIhRiaXN0cm8gc29sb2hhIHByYWd1ZZIBBmJpc3Ryb5oBRENpOURRVWxSUVVOdlpFTm9kSGxqUmpsdlQyczBNMWx0TVhwVk0wNUtaREJ3TlZOSVZuaFZSRTVVVjFod1psZHRZeEFC4AEA-gEFCMIBEDo!16s%2Fg%2F11vzph9zm8?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJjWZ9sayVC0cRypyfz90DtsQ",
							Location: domain.Location{
								Latitude:  50.07385,
								Longitude: 14.4314889,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
						{
							Alias:           "cz-praha-dnister",
							PreviousAliases: nil,
							City:            domain.CzechiaPraha,
							Name:            "Dnister",
							Address:         "Na Moráni 6, 128 00 Praha 2-Nové Město, Česká republika",
							Type:            domain.Restaurant,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/restauracednister?utm_source=qr",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/Dnister/@50.0732496,14.4159963,17z/data=!3m1!4b1!4m6!3m5!1s0x470b94f5d5d2b7d5:0x4eecebf9708c3056!8m2!3d50.0732496!4d14.4159963!16s%2Fg%2F11b6v4fk_x?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ1bfS1fWUC0cRVjCMcPnr7E4",
							Location: domain.Location{
								Latitude:  50.0732496,
								Longitude: 14.4159963,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
						{
							Alias:           "cz-praha-nashi",
							PreviousAliases: nil,
							City:            domain.CzechiaPraha,
							Name:            "NAШІ",
							Address:         "Křižíkova 160/71, 186 00 Praha 8-Karlín, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://nashi.cz/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/nashi_restaurace/",
							FacebookURL:     "",
							GoogleMapsURL:   "https://www.google.com/maps/place/NA%C5%A0I+restaurace/@50.095799,14.4468279,15.35z/data=!4m12!1m2!2m1!1z0J3QkNCo0IYgcmVzdGF1cmFudCBwcmFndWU!3m8!1s0x470b95002f0b29ad:0x67575eb65936dd92!8m2!3d50.0926188!4d14.4510904!9m1!1b1!15sChrQndCQ0KjQhiByZXN0YXVyYW50IHByYWd1ZVocIhrQvdCw0YjRliByZXN0YXVyYW50IHByYWd1ZZIBCnJlc3RhdXJhbnSaASNDaFpEU1VoTk1HOW5TMFZKUTBGblRVUkJNWFZ1YVZOQkVBReABAPoBBQiiARAo!16s%2Fg%2F11wp5cqfs7?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJrSkLLwCVC0cRkt02WbZeV2c",
							Location: domain.Location{
								Latitude:  50.095799,
								Longitude: 14.4468279,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
						{
							Alias:           "cz-praha-the-borsc",
							PreviousAliases: nil,
							City:            domain.CzechiaPraha,
							Name:            "The Boršč",
							Address:         "Dukelských Hrdinů 10, 170 00 Praha 7-Holešovice, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://theborsch.cz/uk",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/the___borsch/?hl=en",
							FacebookURL:     "https://www.facebook.com/TheBorschPrague/",
							GoogleMapsURL:   "https://www.google.com/maps/place/The+Bor%C5%A1%C4%8D/@50.098465,14.433446,3762m/data=!3m1!1e3!4m15!1m8!3m7!1s0x470b94c7babea689:0x57ccaab0ab46b29f!2zRHVrZWxza8O9Y2ggSHJkaW7FryAxMCwgMTcwIDAwIFByYWhhIDctSG9sZcWhb3ZpY2UsINCn0LXRhdC40Y8!3b1!8m2!3d50.0984655!4d14.4334462!16s%2Fg%2F11sjsfrbmr!3m5!1s0x470b952e39f771d1:0x372d3f7ba381c4f8!8m2!3d50.0984655!4d14.4334462!16s%2Fg%2F11x0g10p7t?hl=ru-RU&entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ0XH3OS6VC0cR-MSBo3s_LTc",
							Location: domain.Location{
								Latitude:  50.098465,
								Longitude: 14.433446,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaVranovNadDyji,
					Places: []domain.Place{
						{
							Alias:           "cz-vranov-nad-dyji-bud-laska",
							PreviousAliases: nil,
							City:            domain.CzechiaVranovNadDyji,
							Name:            "Buď Laska",
							Address:         "8. května 92, 671 03 Vranov nad Dyjí, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://www.budlaska.eu/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/budlaskavranov/",
							FacebookURL:     "https://www.facebook.com/budlaskavranov/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Bu%C4%8F+Laska/@48.8964509,15.8127007,17z/data=!4m16!1m9!3m8!1s0x4772adec56937295:0x8162cf21c6ae357!2sBu%C4%8F+Laska!8m2!3d48.8964509!4d15.8127007!9m1!1b1!16s%2Fg%2F11t27jyc5k!3m5!1s0x4772adec56937295:0x8162cf21c6ae357!8m2!3d48.8964509!4d15.8127007!16s%2Fg%2F11t27jyc5k?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJlXKTVuytckcRV-NqHPIsFgg",
							Location: domain.Location{
								Latitude:  48.8964509,
								Longitude: 15.8127007,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
				{
					City: domain.CzechiaUstiNadLabem,
					Places: []domain.Place{
						{
							Alias:           "cz-usti-nad-labem-zlate-karpaty",
							PreviousAliases: nil,
							City:            domain.CzechiaUstiNadLabem,
							Name:            "Restaurace Zlaté Karpaty",
							Address:         "Hornická 2511/10a, 400 11 Ústí nad Labem-centrum, Česká republika",
							Type:            domain.Restaurant,
							URL:             "https://zlatekarpaty.cz/",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "https://www.instagram.com/zlate.karpaty/",
							FacebookURL:     "https://www.facebook.com/p/Restaurace-Zlat%C3%A9-Karpaty-61574925617403/",
							GoogleMapsURL:   "https://www.google.com/maps/place/Restaurace+Zlat%C3%A9+Karpaty/@50.6715865,14.029694,17z/data=!4m16!1m9!3m8!1s0x47099b64006de1f7:0xc88bd350532ba05c!2sRestaurace+Zlat%C3%A9+Karpaty!8m2!3d50.6715865!4d14.029694!9m1!1b1!16s%2Fg%2F11wqh3lj_c!3m5!1s0x47099b64006de1f7:0xc88bd350532ba05c!8m2!3d50.6715865!4d14.029694!16s%2Fg%2F11wqh3lj_c?entry=ttu&g_ep=EgoyMDI2MDYyOC4wIKXMDSoASAFQAw%3D%3D",
							GooglePlaceID:   "ChIJ9-FtAGSbCUcRXKArU1DTi8g",
							Location: domain.Location{
								Latitude:  50.6715865,
								Longitude: 14.029694,
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
							AddedAt:       mustDate("2026-07-01"),
							PromotedUntil: nil,
						},
					},
				},
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
