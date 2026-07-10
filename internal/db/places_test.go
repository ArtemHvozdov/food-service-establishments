package db

import (
	"reflect"
	"testing"
)

// TestPlacesGroupingCountsAndUniqueness перевіряє DoD 1.4: 11 країн, сумарно
// 93 заклади, і жодних дублікатів Place.Alias у межах одного міста.
func TestPlacesGroupingCountsAndUniqueness(t *testing.T) {
	groups := Places()

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
	first := Places()
	second := Places()

	if !reflect.DeepEqual(first, second) {
		t.Fatal("Places() is not deterministic across calls")
	}
}
