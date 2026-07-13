package db

import (
	"fmt"
	"strings"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// dtoToPlace перетворює одну placeDTO (форма JSON) на domain.Place —
// перевалочний крок конвеєра (internal_docs/task_01.md, 1.3). Далі по
// пайплайну DTO не передається (ГОЛОВНИЙ ІНВАРІАНТ).
func dtoToPlace(dto placeDTO) (domain.Place, error) {
	country, ok := countryByCode[dto.Country]
	if !ok {
		return domain.Place{}, fmt.Errorf("db: unresolved country code %q (place id=%q)", dto.Country, dto.ID)
	}

	if strings.TrimSpace(dto.City) == "" {
		return domain.Place{}, fmt.Errorf("db: empty city (place id=%q)", dto.ID)
	}
	cityAlias := slugify(dto.City)
	if cityAlias == "" {
		return domain.Place{}, fmt.Errorf("db: empty city alias for city %q (place id=%q)", dto.City, dto.ID)
	}
	city := domain.City{
		Country: country,
		Name:    dto.City,
		Alias:   cityAlias,
	}

	placeType, ok := placeTypeByString[dto.Type]
	if !ok {
		return domain.Place{}, fmt.Errorf("db: unresolved place type %q (place id=%q)", dto.Type, dto.ID)
	}

	var url string
	if dto.Website != nil {
		url = *dto.Website
	}

	placeAlias := slugify(dto.Name)
	if placeAlias == "" {
		return domain.Place{}, fmt.Errorf("db: empty place alias for name %q (place id=%q)", dto.Name, dto.ID)
	}

	place := domain.Place{
		Alias:         placeAlias,
		City:          city,
		Name:          dto.Name,
		Address:       dto.Address,
		Type:          placeType,
		URL:           url,
		GoogleMapsURL: dto.GoogleMapsURL,
		GooglePlaceID: dto.GooglePlaceID,
		Location: domain.Location{
			Latitude:  dto.Lat,
			Longitude: dto.Lng,
		},
	}

	for _, social := range dto.Socials {
		switch {
		case strings.Contains(social, "facebook.com"):
			place.FacebookURL = social
		case strings.Contains(social, "instagram.com"):
			place.InstagramURL = social
		}
	}

	return place, nil
}
