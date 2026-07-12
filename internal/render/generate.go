package render

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// Generate обходить []domain.CountryPlaceGroup (результат db.Places()) і
// записує все дерево сторінок сайту: головну, сторінки країн (двома
// копіями — internal_docs/task_01.md, 1.8), міст і заведень. Шляхи файлів
// беруться з хелперів 1.6 (render.*FilePath), тому вихідна директорія
// (public/) зафіксована там же.
func Generate(groups []domain.CountryPlaceGroup) error {
	if err := writePage(indexTemplate, groups, IndexFilePath()); err != nil {
		return err
	}

	for _, country := range groups {
		flatFile, indexFile := CountryFilePaths(country.Country)
		if err := writePage(countryTemplate, country, flatFile); err != nil {
			return err
		}
		if err := writePage(countryTemplate, country, indexFile); err != nil {
			return err
		}

		for _, city := range country.Cities {
			if err := writePage(cityTemplate, city, CityFilePath(city.City)); err != nil {
				return err
			}

			for _, place := range city.Places {
				if err := writePage(placeTemplate, place, PlaceFilePath(place)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// writePage виконує шаблон tmpl над data і записує результат у файл path,
// створюючи проміжні директорії за потреби.
func writePage(tmpl *template.Template, data any, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("render: mkdir %s: %w", filepath.Dir(path), err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("render: execute template for %s: %w", path, err)
	}

	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("render: write %s: %w", path, err)
	}

	return nil
}