package db

import "testing"

func validDTO() placeDTO {
	website := "https://example.com"
	return placeDTO{
		ID:            "pl-krakow-example",
		GooglePlaceID: "/g/example",
		Name:          "Example",
		Type:          "restaurant",
		City:          "Kraków",
		Country:       "pl",
		Address:       "Example street",
		Lat:           50.0,
		Lng:           19.0,
		Socials: []string{
			"https://www.facebook.com/example",
			"https://www.instagram.com/example",
		},
		Website:       &website,
		GoogleMapsURL: "https://maps.example.com",
		VerifiedAt:    "2026-06-19",
	}
}

func TestDtoToPlaceResolvesNestedStructs(t *testing.T) {
	place, err := dtoToPlace(validDTO())
	if err != nil {
		t.Fatalf("dtoToPlace() error = %v", err)
	}
	if place.City.Alias == "" || place.City.Name == "" {
		t.Errorf("City not resolved: %+v", place.City)
	}
	if place.City.Country.Alias == "" || place.City.Country.Name == "" {
		t.Errorf("City.Country not resolved: %+v", place.City.Country)
	}
	if place.Type.Alias == "" || place.Type.Name == "" {
		t.Errorf("Type not resolved: %+v", place.Type)
	}
}

func TestDtoToPlaceSocials(t *testing.T) {
	dto := validDTO()
	place, err := dtoToPlace(dto)
	if err != nil {
		t.Fatalf("dtoToPlace() error = %v", err)
	}
	if place.FacebookURL != "https://www.facebook.com/example" {
		t.Errorf("FacebookURL = %q, want facebook URL", place.FacebookURL)
	}
	if place.InstagramURL != "https://www.instagram.com/example" {
		t.Errorf("InstagramURL = %q, want instagram URL", place.InstagramURL)
	}

	dto.Socials = nil
	place, err = dtoToPlace(dto)
	if err != nil {
		t.Fatalf("dtoToPlace() error = %v", err)
	}
	if place.FacebookURL != "" || place.InstagramURL != "" {
		t.Errorf("expected empty social URLs, got FacebookURL=%q InstagramURL=%q", place.FacebookURL, place.InstagramURL)
	}
}

func TestDtoToPlaceNullWebsite(t *testing.T) {
	dto := validDTO()
	dto.Website = nil
	place, err := dtoToPlace(dto)
	if err != nil {
		t.Fatalf("dtoToPlace() error = %v", err)
	}
	if place.URL != "" {
		t.Errorf("URL = %q, want empty string for null website", place.URL)
	}
}

func TestDtoToPlaceUnknownCountry(t *testing.T) {
	dto := validDTO()
	dto.Country = "xx"
	if _, err := dtoToPlace(dto); err == nil {
		t.Fatal("expected error for unknown country code, got nil")
	}
}

func TestDtoToPlaceUnknownType(t *testing.T) {
	dto := validDTO()
	dto.Type = "unknown-type"
	if _, err := dtoToPlace(dto); err == nil {
		t.Fatal("expected error for unknown place type, got nil")
	}
}
