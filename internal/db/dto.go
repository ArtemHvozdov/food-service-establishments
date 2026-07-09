package db

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/ArtemHvozdov/food-service-establishments"
)

// placeDTO — дзеркало одного запису з data/places/*.json (форма даних скрапера).
// Неекспортована: не виходить за межі пакета db (ГОЛОВНИЙ ІНВАРІАНТ, task_01.md).
type placeDTO struct {
	ID            string   `json:"id"`
	GooglePlaceID string   `json:"google_place_id"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	City          string   `json:"city"`
	Country       string   `json:"country"`
	Address       string   `json:"address"`
	Lat           float64  `json:"lat"`
	Lng           float64  `json:"lng"`
	Socials       []string `json:"socials"`
	Website       *string  `json:"website"`
	GoogleMapsURL string   `json:"google_maps_url"`
	VerifiedAt    string   `json:"verified_at"`
}

const placesDir = "data/places"

// loadPlaceDTOs читає всі вшиті data/places/*.json (go:embed) і повертає єдиний
// список placeDTO. Дані первинні: розбіжність з data/countries.json лише
// логується попередженням, парсинг не падає.
func loadPlaceDTOs() ([]placeDTO, error) {
	entries, err := assets.FS.ReadDir(placesDir)
	if err != nil {
		return nil, fmt.Errorf("db: read embedded %s: %w", placesDir, err)
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}
	sort.Strings(names) // детермінований порядок читання файлів

	perCountry := map[string]int{}
	var all []placeDTO
	for _, name := range names {
		if !strings.HasSuffix(name, ".json") {
			continue // пропускаємо не-JSON файли (README, підпапки тощо)
		}

		raw, err := assets.FS.ReadFile(placesDir + "/" + name)
		if err != nil {
			return nil, fmt.Errorf("db: read %s: %w", name, err)
		}
		var dtos []placeDTO
		if err := json.Unmarshal(raw, &dtos); err != nil {
			return nil, fmt.Errorf("db: parse %s: %w", name, err)
		}
		for _, d := range dtos {
			perCountry[d.Country]++
		}
		all = append(all, dtos...)
	}

	warnCountryCountMismatch(perCountry)

	return all, nil
}

// warnCountryCountMismatch звіряє фактичну кількість закладів по країнах з
// data/countries.json. Розбіжність — лише попередження в лог: countries.json
// може відставати від реальних даних (task_01.md, 1.2).
func warnCountryCountMismatch(actual map[string]int) {
	raw, err := assets.FS.ReadFile("data/countries.json")
	if err != nil {
		log.Printf("db: warn: read data/countries.json: %v", err)
		return
	}
	var expected map[string]int
	if err := json.Unmarshal(raw, &expected); err != nil {
		log.Printf("db: warn: parse data/countries.json: %v", err)
		return
	}

	codes := make([]string, 0, len(expected))
	for code := range expected {
		codes = append(codes, code)
	}
	sort.Strings(codes)

	for _, code := range codes {
		if actual[code] != expected[code] {
			log.Printf("db: warn: country %q count mismatch: countries.json=%d, actual=%d", code, expected[code], actual[code])
		}
	}
}
