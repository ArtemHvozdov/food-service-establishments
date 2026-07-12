package render

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ArtemHvozdov/food-service-establishments/internal/db"
)

// generateForTest прогонить реальний конвейєр db.Places() -> Generate() у
// відносну публічну директорію (publicDir), як це робитиме `make generate` з
// кореня репозиторію (1.9), і прибирає за собою після тесту. Working
// directory тестів пакета — сама директорія пакета, тому це не чіпає
// справжній public/ у корені репозиторію.
func generateForTest(t *testing.T) {
	t.Helper()

	groups, err := db.Places()
	if err != nil {
		t.Fatalf("db.Places(): %v", err)
	}

	_ = os.RemoveAll(publicDir)
	t.Cleanup(func() { os.RemoveAll(publicDir) })

	if err := Generate(groups); err != nil {
		t.Fatalf("Generate: %v", err)
	}
}

// TestGenerateWritesFullTree перевіряє дерево public/ після Generate():
// точну кількість файлів кожного рівня і приклади шляхів з DoD 1.8.
func TestGenerateWritesFullTree(t *testing.T) {
	generateForTest(t)

	var indexCount, countryFlatCount, countryIndexCount, cityCount, placeCount, total int

	err := filepath.WalkDir(publicDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		total++

		rel, err := filepath.Rel(publicDir, path)
		if err != nil {
			return err
		}
		segments := strings.Split(rel, string(filepath.Separator))

		switch len(segments) {
		case 1: // public/index.html або public/[country].html
			if segments[0] == "index.html" {
				indexCount++
			} else {
				countryFlatCount++
			}
		case 2: // public/[country]/index.html або public/[country]/[city].html
			if segments[1] == "index.html" {
				countryIndexCount++
			} else {
				cityCount++
			}
		case 3: // public/[country]/[city]/[place].html
			placeCount++
		default:
			t.Errorf("неочікувана глибина файлу: %s", path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDir: %v", err)
	}

	if indexCount != 1 {
		t.Errorf("index.html: очікувалось 1, отримано %d", indexCount)
	}
	if countryFlatCount != 11 {
		t.Errorf("public/[country].html: очікувалось 11, отримано %d", countryFlatCount)
	}
	if countryIndexCount != 11 {
		t.Errorf("public/[country]/index.html: очікувалось 11, отримано %d", countryIndexCount)
	}
	if cityCount != 57 {
		t.Errorf("public/[country]/[city].html: очікувалось 57, отримано %d", cityCount)
	}
	if placeCount != 93 {
		t.Errorf("public/[country]/[city]/[place].html: очікувалось 93, отримано %d", placeCount)
	}
	if total != 173 {
		t.Errorf("всього html-файлів: очікувалось 173, отримано %d", total)
	}

	for _, p := range []string{
		filepath.Join(publicDir, "poland.html"),
		filepath.Join(publicDir, "poland", "index.html"),
		filepath.Join(publicDir, "poland", "krakow.html"),
	} {
		if _, err := os.Stat(p); err != nil {
			t.Errorf("очікувався файл %s: %v", p, err)
		}
	}

	placeEntries, err := os.ReadDir(filepath.Join(publicDir, "poland", "krakow"))
	if err != nil {
		t.Fatalf("public/poland/krakow/: %v", err)
	}
	if len(placeEntries) == 0 {
		t.Errorf("очікувався хоча б один файл заведення у public/poland/krakow/")
	}
}

// TestGenerateHasNoDeadLinks перевіряє, що кожен внутрішній href="/..."
// у згенерованих сторінках веде на реально існуючий файл.
func TestGenerateHasNoDeadLinks(t *testing.T) {
	generateForTest(t)

	hrefRe := regexp.MustCompile(`href="(/[^"]*)"`)
	checked := 0

	err := filepath.WalkDir(publicDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		for _, m := range hrefRe.FindAllStringSubmatch(string(content), -1) {
			href := m[1]
			target := hrefToFilePath(href)
			if _, statErr := os.Stat(target); statErr != nil {
				t.Errorf("%s: битий href %q -> %s: %v", path, href, target, statErr)
			}
			checked++
		}
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDir: %v", err)
	}
	if checked == 0 {
		t.Fatal("жодного href не знайдено — перевірка не спрацювала")
	}
}

// hrefToFilePath перетворює кореневий URL (як повертають хелпери 1.6) на
// шлях файлу в публічній директорії.
func hrefToFilePath(href string) string {
	if href == "/" {
		return filepath.Join(publicDir, "index.html")
	}
	return filepath.Join(publicDir, filepath.FromSlash(strings.TrimPrefix(href, "/")))
}
