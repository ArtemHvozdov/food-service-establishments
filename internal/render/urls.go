// Package render містить хелпери, спільні для генерації сторінок: шляхи
// файлів у public/ та відповідні кореневі URL (internal_docs/task_01.md, 1.6),
// а згодом — шаблони (1.7) і сам обхід/запис дерева (1.8).
// Працює виключно з domain.*: DTO-шар (internal/db) сюди не потрапляє (ГЛАВНЫЙ ИНВАРИАНТ).
package render

import (
	"path"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

// publicDir — коренева директорія згенерованого сайту.
const publicDir = "public"

// IndexFilePath повертає шлях файлу головної сторінки.
func IndexFilePath() string {
	return path.Join(publicDir, "index.html")
}

// IndexURL повертає кореневий URL головної сторінки.
func IndexURL() string {
	return "/"
}

// CountryFilePaths повертає ДВА шляхи файлу сторінки країни: плоский
// public/[country].html і вкладений public/[country]/index.html. Обидва
// файли пишуться однаковим шаблоном (1.8) — це навмисне дублювання заради
// сумісності з різними правилами маршрутизації статичних хостингів.
func CountryFilePaths(country domain.Country) (flatFile, indexFile string) {
	flatFile = path.Join(publicDir, country.Alias+".html")
	indexFile = path.Join(publicDir, country.Alias, "index.html")
	return flatFile, indexFile
}

// CountryURL повертає кореневий URL сторінки країни. Він один для обох копій
// файлу (canonical завжди вказує на плоский варіант /[country].html).
func CountryURL(country domain.Country) string {
	return "/" + country.Alias + ".html"
}

// CityFilePath повертає шлях файлу сторінки міста.
func CityFilePath(city domain.City) string {
	return path.Join(publicDir, city.Country.Alias, city.Alias+".html")
}

// CityURL повертає кореневий URL сторінки міста.
func CityURL(city domain.City) string {
	return "/" + city.Country.Alias + "/" + city.Alias + ".html"
}

// PlaceFilePath повертає шлях файлу сторінки заведення.
func PlaceFilePath(place domain.Place) string {
	return path.Join(publicDir, place.City.Country.Alias, place.City.Alias, place.Alias+".html")
}

// PlaceURL повертає кореневий URL сторінки заведення.
func PlaceURL(place domain.Place) string {
	return "/" + place.City.Country.Alias + "/" + place.City.Alias + "/" + place.Alias + ".html"
}
