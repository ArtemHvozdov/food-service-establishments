package db

import (
	"strings"
	"testing"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// TestValidatePassesOnRealData перевіряє DoD 1.5: на поточних коректних
// даних Validate не повертає помилок.
func TestValidatePassesOnRealData(t *testing.T) {
	groups, err := Places()
	if err != nil {
		t.Fatalf("Places() returned error on real data: %v", err)
	}
	if err := Validate(groups); err != nil {
		t.Fatalf("Validate(Places()) returned error on real data: %v", err)
	}
}

// validGroup — мінімальний коректний []domain.CountryPlaceGroup для
// негативних тестів нижче: кожен тест псує рівно один інваріант.
func validGroup() []domain.CountryPlaceGroup {
	return []domain.CountryPlaceGroup{
		{
			Country: domain.Poland,
			Cities: []domain.CityPlaceGroup{
				{
					City: domain.City{Country: domain.Poland, Alias: "krakow", Name: "Краків"},
					Places: []domain.Place{
						{Name: "Borshch", Alias: "borshch", URL: "https://example.com"},
					},
				},
			},
		},
	}
}

func TestValidateNoOutgoingLinkFails(t *testing.T) {
	groups := validGroup()
	groups[0].Cities[0].Places[0].URL = ""

	err := Validate(groups)
	if err == nil {
		t.Fatal("expected error for place without outgoing link, got nil")
	}
	if !strings.Contains(err.Error(), "no outgoing link") {
		t.Fatalf("error %q does not mention missing link", err.Error())
	}
}

func TestValidateEmptyPlaceNameFails(t *testing.T) {
	groups := validGroup()
	groups[0].Cities[0].Places[0].Name = ""

	err := Validate(groups)
	if err == nil {
		t.Fatal("expected error for empty Place.Name, got nil")
	}
	if !strings.Contains(err.Error(), "empty Place.Name") {
		t.Fatalf("error %q does not mention empty name", err.Error())
	}
}

func TestValidateDuplicatePlaceAliasInCityFails(t *testing.T) {
	groups := validGroup()
	groups[0].Cities[0].Places = append(groups[0].Cities[0].Places, domain.Place{
		Name: "Borshch 2", Alias: "borshch", URL: "https://example.com",
	})

	err := Validate(groups)
	if err == nil {
		t.Fatal("expected error for duplicate Place.Alias in city, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate Place.Alias") {
		t.Fatalf("error %q does not mention duplicate alias", err.Error())
	}
}
