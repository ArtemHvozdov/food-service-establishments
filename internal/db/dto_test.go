package db

import "testing"

func TestLoadPlaceDTOsCount(t *testing.T) {
	dtos, err := loadPlaceDTOs()
	if err != nil {
		t.Fatalf("loadPlaceDTOs() error = %v", err)
	}
	if got, want := len(dtos), 93; got != want {
		t.Fatalf("loadPlaceDTOs() returned %d DTOs, want %d", got, want)
	}

	perCountry := map[string]int{}
	for _, d := range dtos {
		perCountry[d.Country]++
	}
	// countries.json: pl:17 de:22 ro:6 es:7 cz:13 tr:3 pt:8 bg:3 cy:3 sk:6 hu:4.
	// cz фактично містить 14 записів (countries.json відстає) — це очікувана
	// розбіжність, не помилка (task_01.md, 1.2).
	want := map[string]int{
		"pl": 17, "de": 22, "ro": 6, "es": 7, "cz": 14,
		"tr": 3, "pt": 8, "bg": 3, "cy": 3, "sk": 6, "hu": 4,
	}
	for code, wantCount := range want {
		if perCountry[code] != wantCount {
			t.Errorf("country %q: got %d places, want %d", code, perCountry[code], wantCount)
		}
	}
}

func TestLoadPlaceDTOsNullableWebsite(t *testing.T) {
	dtos, err := loadPlaceDTOs()
	if err != nil {
		t.Fatalf("loadPlaceDTOs() error = %v", err)
	}
	foundNull := false
	for _, d := range dtos {
		if d.Website == nil {
			foundNull = true
			break
		}
	}
	if !foundNull {
		t.Fatalf("expected at least one DTO with Website == nil (website: null in JSON)")
	}
}
