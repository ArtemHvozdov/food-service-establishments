package db

import (
	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

func Places() []domain.CountryPlaceGroup {
	krakow := domain.City{
		Country: domain.Poland,
		Alias:   "krakow",
		Name:    "Краків",
	}

	return []domain.CountryPlaceGroup{
		{
			Country: domain.Poland,
			Cities: []domain.CityPlaceGroup{
				{
					City: krakow,
					Places: []domain.Place{
						{
							Alias:           "",
							PreviousAliases: nil,
							City:            krakow,
							Name:            "",
							Address:         "",
							Type:            domain.Restaurant,
							URL:             "",
							MenuURL:         "",
							PhotoURL:        "",
							InstagramURL:    "",
							GoogleMapsURL:   "",
							GooglePlaceID:   "",
							Location:        domain.Location{},
							FoundingYear:    0,
							Description:     "",
							Dishes:          nil,
							BorschPrice:     "",
							LargePortions:   false,
							HasDelivery:     false,
							DonatesToZSU:    nil,
							Articles:        nil,
							YoutubeVideos:   nil,
							AddedAt:         mustDate("2026-07-01"),
							PromotedUntil:   nil,
						},
					},
				},
			},
		},
	}
}
