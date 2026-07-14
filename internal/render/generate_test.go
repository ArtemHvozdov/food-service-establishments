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

// generateForTest прогонить реальний конвейєр db.PlacesOld() -> Generate() у
// ізольовану тимчасову директорію (t.TempDir()), яку Go сам прибирає після
// тесту. Не чіпає справжній public/ у корені репозиторію.
func generateForTest(t *testing.T) string {
	t.Helper()

	groups, err := db.PlacesOld()
	if err != nil {
		t.Fatalf("db.PlacesOld(): %v", err)
	}

	dir := t.TempDir()
	if err := Generate(groups, dir); err != nil {
		t.Fatalf("Generate: %v", err)
	}
	return dir
}

// TestGenerateWritesFullTree перевіряє дерево public/ після Generate():
// точну кількість файлів кожного рівня і приклади шляхів з DoD 1.8.
func TestGenerateWritesFullTree(t *testing.T) {
	dir := generateForTest(t)

	var indexCount, countryFlatCount, countryIndexCount, cityCount, placeCount, total int

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		total++

		rel, err := filepath.Rel(dir, path)
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
	if countryFlatCount != 10 {
		t.Errorf("public/[country].html: очікувалось 10, отримано %d", countryFlatCount)
	}
	if countryIndexCount != 10 {
		t.Errorf("public/[country]/index.html: очікувалось 10, отримано %d", countryIndexCount)
	}
	if cityCount != 55 {
		t.Errorf("public/[country]/[city].html: очікувалось 55, отримано %d", cityCount)
	}
	if placeCount != 90 {
		t.Errorf("public/[country]/[city]/[place].html: очікувалось 90, отримано %d", placeCount)
	}
	if total != 166 {
		t.Errorf("всього html-файлів: очікувалось 166, отримано %d", total)
	}

	for _, p := range []string{
		filepath.Join(dir, "poland.html"),
		filepath.Join(dir, "poland", "index.html"),
		filepath.Join(dir, "poland", "krakow.html"),
	} {
		if _, err := os.Stat(p); err != nil {
			t.Errorf("очікувався файл %s: %v", p, err)
		}
	}

	placeEntries, err := os.ReadDir(filepath.Join(dir, "poland", "krakow"))
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
	dir := generateForTest(t)

	hrefRe := regexp.MustCompile(`href="(/[^"]*)"`)
	checked := 0

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
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
			target := hrefToFilePath(dir, href)
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
// шлях файлу всередині вихідної директорії dir.
func hrefToFilePath(dir, href string) string {
	if href == "/" {
		return filepath.Join(dir, "index.html")
	}
	return filepath.Join(dir, filepath.FromSlash(strings.TrimPrefix(href, "/")))
}
