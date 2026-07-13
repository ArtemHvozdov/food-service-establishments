// Package render містить хелпери, спільні для генерації сторінок: шляхи
// файлів у public/ та відповідні кореневі URL (internal_docs/task_01.md, 1.6),
// а згодом — шаблони (1.7) і сам обхід/запис дерева (1.8).
// Працює виключно з domain.*: DTO-шар (internal/db) сюди не потрапляє (ГЛАВНЫЙ ИНВАРИАНТ).
package render

import (
	"path/filepath"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// IndexFilePath повертає шлях файлу головної сторінки всередині outputDir.
func IndexFilePath(outputDir string) string {
	return filepath.Join(outputDir, "index.html")
}

// IndexURL повертає кореневий URL головної сторінки.
func IndexURL() string {
	return "/"
}

// CountryFilePaths повертає ДВА шляхи файлу сторінки країни всередині
// outputDir: плоский [outputDir]/[country].html і вкладений
// [outputDir]/[country]/index.html. Обидва файли пишуться однаковим шаблоном
// (1.8) — це навмисне дублювання заради сумісності з різними правилами
// маршрутизації статичних хостингів.
func CountryFilePaths(outputDir string, country domain.Country) (flatFile, indexFile string) {
	flatFile = filepath.Join(outputDir, country.Alias+".html")
	indexFile = filepath.Join(outputDir, country.Alias, "index.html")
	return flatFile, indexFile
}

// CountryURL повертає кореневий URL сторінки країни. Він один для обох копій
// файлу (canonical завжди вказує на плоский варіант /[country].html).
func CountryURL(country domain.Country) string {
	return "/" + country.Alias + ".html"
}

// CityFilePath повертає шлях файлу сторінки міста всередині outputDir.
func CityFilePath(outputDir string, city domain.City) string {
	return filepath.Join(outputDir, city.Country.Alias, city.Alias+".html")
}

// CityURL повертає кореневий URL сторінки міста.
func CityURL(city domain.City) string {
	return "/" + city.Country.Alias + "/" + city.Alias + ".html"
}

// PlaceFilePath повертає шлях файлу сторінки заведення всередині outputDir.
func PlaceFilePath(outputDir string, place domain.Place) string {
	return filepath.Join(outputDir, place.City.Country.Alias, place.City.Alias, place.Alias+".html")
}

// PlaceURL повертає кореневий URL сторінки заведення.
func PlaceURL(place domain.Place) string {
	return "/" + place.City.Country.Alias + "/" + place.City.Alias + "/" + place.Alias + ".html"
}
