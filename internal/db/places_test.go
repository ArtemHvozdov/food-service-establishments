package db

import (
	"reflect"
	"testing"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// TestPlacesGroupingCountsAndUniqueness перевіряє DoD 1.4: 11 країн, сумарно
// 93 заклади, і жодних дублікатів Place.Alias у межах одного міста.
func TestPlacesGroupingCountsAndUniqueness(t *testing.T) {
	groups, err := Places()
	if err != nil {
		t.Fatalf("Places() error = %v", err)
	}

	if got, want := len(groups), 11; got != want {
		t.Errorf("len(groups) = %d, want %d", got, want)
	}

	total := 0
	for _, country := range groups {
		for _, city := range country.Cities {
			seen := map[string]bool{}
			for _, place := range city.Places {
				if seen[place.Alias] {
					t.Errorf("duplicate Place.Alias %q in city %q/%q", place.Alias, country.Country.Alias, city.City.Alias)
				}
				seen[place.Alias] = true
				total++
			}
		}
	}

	if got, want := total, 93; got != want {
		t.Errorf("total places = %d, want %d", got, want)
	}
}

// TestPlacesDeterministic перевіряє, що повторний виклик Places() дає
// ідентичний порядок (DoD 1.4: детермінізм між білдами).
func TestPlacesDeterministic(t *testing.T) {
	first, err := Places()
	if err != nil {
		t.Fatalf("First Places() error = %v", err)
	}
	second, err := Places()
	if err != nil {
		t.Fatalf("Second Places() error = %v", err)
	}

	if !reflect.DeepEqual(first, second) {
		t.Fatal("Places() is not deterministic across calls")
	}
}

// TestDedupeAliasesAvoidsNaturalCollision — крайовий випадок 1.4.1: два
// заведення "chernomorka" плюс "натуральне" chernomorka-2 не повинні дати
// дублікат після дедупу (другий "chernomorka" мусить перескочити на -3).
func TestDedupeAliasesAvoidsNaturalCollision(t *testing.T) {
	places := []domain.Place{
		{Alias: "chernomorka"},
		{Alias: "chernomorka"},
		{Alias: "chernomorka-2"},
	}

	dedupeAliases(places)

	seen := map[string]bool{}
	for _, p := range places {
		if seen[p.Alias] {
			t.Fatalf("duplicate alias %q after dedupeAliases: %+v", p.Alias, places)
		}
		seen[p.Alias] = true
	}
}

// TestGroupPlacesSuffixAssignmentStableAcrossInputOrder — DoD 1.4.1: при
// повному збігу Name/Alias призначення суфікса -2 має триматися стабільного
// тай-брейка (GooglePlaceID), а не позиції запису у вхідному файлі.
func TestGroupPlacesSuffixAssignmentStableAcrossInputOrder(t *testing.T) {
	city := domain.City{Country: domain.Slovakia, Alias: "bratislava", Name: "Братислава"}

	a := domain.Place{City: city, Name: "Chernomorka", Alias: "chernomorka", GooglePlaceID: "/g/aaa"}
	b := domain.Place{City: city, Name: "Chernomorka", Alias: "chernomorka", GooglePlaceID: "/g/bbb"}

	aliasByGooglePlaceID := func(groups []domain.CountryPlaceGroup) map[string]string {
		result := map[string]string{}
		for _, country := range groups {
			for _, city := range country.Cities {
				for _, p := range city.Places {
					result[p.GooglePlaceID] = p.Alias
				}
			}
		}
		return result
	}

	forward := aliasByGooglePlaceID(groupPlaces([]domain.Place{a, b}))
	backward := aliasByGooglePlaceID(groupPlaces([]domain.Place{b, a}))

	if !reflect.DeepEqual(forward, backward) {
		t.Fatalf("suffix assignment depends on input order: forward=%v backward=%v", forward, backward)
	}
}
