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
	return "/" + country.Alias + "/"
}

// CityFilePath повертає шлях файлу сторінки міста всередині outputDir.
func CityFilePath(outputDir string, city domain.City) (flatFile, indexFile string) {
	flatFile = filepath.Join(outputDir, city.Country.Alias, city.Alias+".html")
	indexFile = filepath.Join(outputDir, city.Country.Alias, city.Alias, "index.html")
	return flatFile, indexFile
}

// CityURL повертає кореневий URL сторінки міста.
func CityURL(city domain.City) string {
	return "/" + city.Country.Alias + "/" + city.Alias + "/"
}

// PlaceFilePath повертає шлях файлу сторінки заведення всередині outputDir.
func PlaceFilePath(outputDir string, place domain.Place) (flatFile, indexFile string) {
	flatFile = filepath.Join(outputDir, place.City.Country.Alias, place.City.Alias, place.Alias+".html")
	indexFile = filepath.Join(outputDir, place.City.Country.Alias, place.City.Alias, place.Alias, "index.html")
	return flatFile, indexFile
}

// PlaceURL повертає кореневий URL сторінки заведення.
func PlaceURL(place domain.Place) string {
	return "/" + place.City.Country.Alias + "/" + place.City.Alias + "/" + place.Alias + "/"
}
