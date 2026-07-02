package db

import (
	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

func Places() []domain.CountryPlaceGroup {
	return []domain.CountryPlaceGroup{
		{
			Country: domain.Poland,
			Cities: []domain.CityPlaceGroup{
				{
					City: domain.Krakow,
					Places: []domain.Place{
						{
							Alias:           "",
							PreviousAliases: nil,
							City:            domain.Krakow,
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
