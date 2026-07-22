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
// записує все дерево сторінок сайту в outputDir: головну, сторінки країн
// (двома копіями — internal_docs/task_01.md, 1.8), міст і заведень. Шляхи
// файлів беруться з хелперів 1.6 (render.*FilePath).
func Generate(groups []domain.CountryPlaceGroup, outputDir string) error {
	if err := writePage(indexTemplate, groups, IndexFilePath(outputDir)); err != nil {
		return err
	}

	for _, countryGroup := range groups {
		countryFlatFile, countryIndexFile := CountryFilePaths(outputDir, countryGroup.Country)
		if err := writePage(countryTemplate, countryGroup, countryFlatFile); err != nil {
			return err
		}
		if err := writePage(countryTemplate, countryGroup, countryIndexFile); err != nil {
			return err
		}

		for _, cityGroup := range countryGroup.Cities {
			cityFlatFile, cityIndexFile := CityFilePath(outputDir, cityGroup.City)
			if err := writePage(cityTemplate, cityGroup, cityFlatFile); err != nil {
				return err
			}
			if err := writePage(cityTemplate, cityGroup, cityIndexFile); err != nil {
				return err
			}

			for _, place := range cityGroup.Places {
				placeFlatFile, placeIndexFile := PlaceFilePath(outputDir, place)

				if err := writePage(placeTemplate, place, placeFlatFile); err != nil {
					return err
				}
				if err := writePage(placeTemplate, place, placeIndexFile); err != nil {
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
