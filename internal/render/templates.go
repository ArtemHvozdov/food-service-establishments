package render

import (
	"embed"
	"fmt"
	"html/template"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

//go:embed templates/layout.html templates/place.html templates/city.html templates/country.html templates/index.html
var templateFS embed.FS

// funcMap — функції, доступні шаблонам. Посилання (href) і head-метадані
// (title/description/canonical) обчислюються тут через хелпери 1.6, а не
// збираються рядками прямо в шаблоні (internal_docs/task_01.md, 1.7).
var funcMap = template.FuncMap{
	"placeURL":        PlaceURL,
	"cityURL":         CityURL,
	"countryURL":      CountryURL,
	"pageTitle":       pageTitle,
	"pageDescription": pageDescription,
	"canonicalOf":     canonicalOf,
}

// newPageTemplate компілює спільний layout ("layout.html") разом з одним
// файлом контенту сторінки (той визначає {{define "content"}}). Кожен тип
// сторінки має свій незалежний *template.Template, тому що всі файли
// контенту визначають блок з однаковим ім'ям "content" — розбір їх усіх
// разом в один набір призвів би до взаємного перезапису.
func newPageTemplate(contentFile string) *template.Template {
	tmpl := template.Must(template.New("layout").Funcs(funcMap).ParseFS(templateFS, "templates/layout.html"))
	return template.Must(tmpl.ParseFS(templateFS, contentFile))
}

var (
	// placeTemplate рендерить сторінку заведення. Data-контекст: domain.Place.
	placeTemplate = newPageTemplate("templates/place.html")
	// cityTemplate рендерить сторінку міста. Data-контекст: domain.CityPlaceGroup.
	cityTemplate = newPageTemplate("templates/city.html")
	// countryTemplate рендерить сторінку країни. Data-контекст: domain.CountryPlaceGroup.
	countryTemplate = newPageTemplate("templates/country.html")
	// indexTemplate рендерить головну сторінку. Data-контекст: []domain.CountryPlaceGroup.
	indexTemplate = newPageTemplate("templates/index.html")
)

// pageTitle будує <title> сторінки за типом даних, переданих у шаблон.
// Аргумент — точно той-таки доменний контекст, що й дані сторінки: DTO сюди
// не потрапляє.
func pageTitle(v any) string {
	const suffix = "Українська кухня за кордоном"
	switch data := v.(type) {
	case domain.Place:
		return fmt.Sprintf("%s — %s, %s | %s", data.Name, data.City.Name, data.City.Country.Name, suffix)
	case domain.CityPlaceGroup:
		return fmt.Sprintf("Заклади української кухні у %s, %s | %s", data.City.Name, data.City.Country.Name, suffix)
	case domain.CountryPlaceGroup:
		return fmt.Sprintf("Заклади української кухні у %s | %s", data.Country.Name, suffix)
	case []domain.CountryPlaceGroup:
		return "Каталог закладів " + suffix
	default:
		return suffix
	}
}

// pageDescription будує <meta name="description"> сторінки за типом даних.
func pageDescription(v any) string {
	switch data := v.(type) {
	case domain.Place:
		return fmt.Sprintf("%s — заклад української кухні у %s, %s. Адреса, сайт і соціальні мережі.", data.Name, data.City.Name, data.City.Country.Name)
	case domain.CityPlaceGroup:
		return fmt.Sprintf("Список закладів української кухні у %s, %s: ресторани, кафе, бістро, бари, пекарні.", data.City.Name, data.City.Country.Name)
	case domain.CountryPlaceGroup:
		return fmt.Sprintf("Список закладів української кухні у %s за містами.", data.Country.Name)
	case []domain.CountryPlaceGroup:
		return "Каталог закладів української кухні за кордоном: ресторани, кафе, бістро, бари та пекарні за країнами й містами."
	default:
		return ""
	}
}

// canonicalOf повертає кореневий canonical URL сторінки за типом даних,
// побудований хелперами 1.6.
func canonicalOf(v any) string {
	switch data := v.(type) {
	case domain.Place:
		return PlaceURL(data)
	case domain.CityPlaceGroup:
		return CityURL(data.City)
	case domain.CountryPlaceGroup:
		return CountryURL(data.Country)
	case []domain.CountryPlaceGroup:
		return IndexURL()
	default:
		return ""
	}
}
