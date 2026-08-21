package render

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/ArtemHvozdov/food-service-establishments/internal/domain"
)

//go:embed templates/layout.html templates/partials.html templates/place.html templates/city.html templates/country.html templates/index.html
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
	"breadcrumbs":     breadcrumbs,
	"upLink":          upLink,
	"absURL":          absURL,
	"ogImage":         ogImage,

	"breadcrumbListJSONLD": breadcrumbListJSONLD,
	"itemListJSONLD":       itemListJSONLD,
	"placeJSONLD":          placeJSONLD,
}

// newPageTemplate компілює спільний layout ("layout.html") і спільні partial'и
// навігації ("partials.html") разом з одним файлом контенту сторінки (той
// визначає {{define "content"}}). Кожен тип сторінки має свій незалежний
// *template.Template, тому що всі файли контенту визначають блок з однаковим
// ім'ям "content" — розбір їх усіх разом в один набір призвів би до
// взаємного перезапису.
func newPageTemplate(contentFile string) *template.Template {
	tmpl := template.Must(template.New("layout").Funcs(funcMap).ParseFS(templateFS, "templates/layout.html", "templates/partials.html"))
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

// ogImage повертає абсолютний URL og:image (задача 3.4, доопрацьовано в 3.6,
// design.md §6.1) — тільки для сторінки заведення й тільки коли
// Place.PhotoURL заповнений. PhotoURL зберігається як кореневий шлях
// ("/photos/...", Validate це перевіряє), так само як і URL з
// canonicalOf/breadcrumbs, тож тут так само потрібен absURL — інакше
// og:image був би відносним, що Google Rich Results не резолвить. Поки фото
// немає в даних, повертає "" — layout.html тоді не рендерить тег og:image
// взагалі (порожній og:image гірший за його відсутність); саме тому перевірка
// на порожній PhotoURL — до виклику absURL, бо absURL("") дав би хибний
// непорожній baseURL.
func ogImage(v any) string {
	place, ok := v.(domain.Place)
	if !ok || place.Photo.Path == "" {
		return ""
	}
	return absURL(place.Photo.Path)
}

// crumb — один елемент хлібної крихти (design.md §4.1). Порожній URL означає
// поточну сторінку — вона рендериться текстом, не посиланням.
type crumb struct {
	Name string
	URL  string
}

// breadcrumbs будує ланцюжок хлібних крихт за типом даних сторінки.
// Архітектурне рішення (задача 3.2): окремий view-model рівня render, що
// обгортав би дані сторінки предками, тут не потрібен — Place.City і
// City.Country вже дають прямий доступ до всіх предків, тож helper просто
// повторює прийом pageTitle/pageDescription/canonicalOf: перемикання за
// типом контексту, переданого в шаблон. На головній (`[]domain.CountryPlaceGroup`)
// крихт немає взагалі (design.md §4.1) — nil.
func breadcrumbs(v any) []crumb {
	switch data := v.(type) {
	case domain.Place:
		return []crumb{
			{Name: "Головна", URL: IndexURL()},
			{Name: data.City.Country.Name, URL: CountryURL(data.City.Country)},
			{Name: data.City.Name, URL: CityURL(data.City)},
			{Name: data.Name},
		}
	case domain.CityPlaceGroup:
		return []crumb{
			{Name: "Головна", URL: IndexURL()},
			{Name: data.City.Country.Name, URL: CountryURL(data.City.Country)},
			{Name: data.City.Name},
		}
	case domain.CountryPlaceGroup:
		return []crumb{
			{Name: "Головна", URL: IndexURL()},
			{Name: data.Country.Name},
		}
	default:
		return nil
	}
}

// backLink — посилання «вгору на рівень» унизу не-головної сторінки (design.md §4.2).
type backLink struct {
	Text string
	URL  string
}

// upLink повертає посилання на батьківський рівень за типом даних сторінки,
// або nil для головної (там посилання вгору немає — нікуди підійматися).
func upLink(v any) *backLink {
	switch data := v.(type) {
	case domain.Place:
		return &backLink{Text: "← Усі заклади в " + data.City.Name, URL: CityURL(data.City)}
	case domain.CityPlaceGroup:
		return &backLink{Text: "← Усі міста в " + data.City.Country.Name, URL: CountryURL(data.City.Country)}
	case domain.CountryPlaceGroup:
		return &backLink{Text: "← Усі країни", URL: IndexURL()}
	default:
		return nil
	}
}

// Далі — JSON-LD (schema.org) розмітка для SEO/AI-SEO (задача 3.3, design.md).
// Усі *JSONLD-функції повертають template.JS: html/template визнає цей тип
// довіреним і не застосовує повторне екранування до вже коректного JSON,
// зібраного через encoding/json (той сам екранує "<"/">"/"&", тож розрив
// "</script>" усередині значень неможливий).

// marshalJSONLD маршалить v у JSON і повертає його як template.JS для вставки
// в <script type="application/ld+json">. Помилка маршалінгу тут неможлива —
// усі *JSONLD-структури складені з рядків/чисел/зрізів таких самих структур.
func marshalJSONLD(v any) template.JS {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return template.JS(b)
}

// breadcrumbListItem — один елемент JSON-LD BreadcrumbList.
type breadcrumbListItem struct {
	Type     string `json:"@type"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item,omitempty"`
}

// breadcrumbListJSONLD будує JSON-LD BreadcrumbList на основі тих самих даних
// про предків, що йдуть і в HTML-крихти (design.md §4.1, breadcrumbs вище).
// На головній breadcrumbs(v) повертає nil — тоді й тут порожньо.
func breadcrumbListJSONLD(v any) template.JS {
	crumbs := breadcrumbs(v)
	if len(crumbs) == 0 {
		return ""
	}

	items := make([]breadcrumbListItem, len(crumbs))
	for i, c := range crumbs {
		url := c.URL
		if url == "" {
			// Останній елемент — поточна сторінка (в HTML-крихтах без
			// посилання); у JSON-LD посилаємось на її ж canonical.
			url = canonicalOf(v)
		}
		// JSON-LD item — абсолютний URL (задача 3.3.2): Google Rich Results
		// не резолвить кореневі відносні шляхи.
		items[i] = breadcrumbListItem{Type: "ListItem", Position: i + 1, Name: c.Name, Item: absURL(url)}
	}

	return marshalJSONLD(struct {
		Context         string               `json:"@context"`
		Type            string               `json:"@type"`
		ItemListElement []breadcrumbListItem `json:"itemListElement"`
	}{
		Context:         "https://schema.org",
		Type:            "BreadcrumbList",
		ItemListElement: items,
	})
}

// listItem — один елемент JSON-LD ItemList (сторінки-списки країни/міста).
type listItem struct {
	Type     string `json:"@type"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	URL      string `json:"url"`
}

// itemListJSONLD будує JSON-LD ItemList для сторінок-списків: країна → міста,
// місто → заклади (design.md §4.3). Для інших рівнів (головна, заклад) —
// порожньо: на головній список країн не описаний окремим інвентарем задачі
// 3.3, а сторінка заклада описується placeJSONLD.
func itemListJSONLD(v any) template.JS {
	var items []listItem
	switch data := v.(type) {
	case domain.CountryPlaceGroup:
		for i, cg := range data.Cities {
			items = append(items, listItem{Type: "ListItem", Position: i + 1, Name: cg.City.Name, URL: absURL(CityURL(cg.City))})
		}
	case domain.CityPlaceGroup:
		for i, p := range data.Places {
			items = append(items, listItem{Type: "ListItem", Position: i + 1, Name: p.Name, URL: absURL(PlaceURL(p))})
		}
	default:
		return ""
	}
	if len(items) == 0 {
		return ""
	}

	return marshalJSONLD(struct {
		Context         string     `json:"@context"`
		Type            string     `json:"@type"`
		ItemListElement []listItem `json:"itemListElement"`
	}{
		Context:         "https://schema.org",
		Type:            "ItemList",
		ItemListElement: items,
	})
}

// schemaOrgPlaceType мапить domain.PlaceType на найближчий тип schema.org
// (design.md/DoD 3.3: "Restaurant/FoodEstablishment — маппінг з PlaceType").
// Для типів без точного відповідника в schema.org (bistro) — узагальнений
// FoodEstablishment, а не вигаданий тип.
func schemaOrgPlaceType(t domain.PlaceType) string {
	switch t.Alias {
	case domain.Restaurant.Alias:
		return "Restaurant"
	case domain.Cafe.Alias:
		return "CafeOrCoffeeShop"
	case domain.Bar.Alias:
		return "BarOrPub"
	case domain.Bakery.Alias:
		return "Bakery"
	default:
		return "FoodEstablishment"
	}
}

// cityLatinName наближено відновлює латинську назву міста для JSON-LD
// addressLocality (задача 3.3.3): окремого латинського поля назви міста в
// domain.City немає, тому за основу береться Alias (завжди латиницею,
// напр. "karlovy-vary") і кожна частина через "-" переводиться в Title Case
// та з'єднується пробілом ("Karlovy Vary"). Це наближення (може не
// збігатися з офіційною латинською формою на 100%), обране як компроміс
// без розширення domain.City заради єдиного SEO-поля.
func cityLatinName(alias string) string {
	parts := strings.Split(alias, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}

// jsonLDPostalAddress — schema.org PostalAddress, вбудований у placeJSONLD.
type jsonLDPostalAddress struct {
	Type            string `json:"@type"`
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	AddressCountry  string `json:"addressCountry"`
}

// jsonLDGeoCoordinates — schema.org GeoCoordinates, вбудований у placeJSONLD.
type jsonLDGeoCoordinates struct {
	Type      string  `json:"@type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// jsonLDPlace — JSON-LD Restaurant/FoodEstablishment (schema.org) заведення.
type jsonLDPlace struct {
	Context string               `json:"@context"`
	Type    string               `json:"@type"`
	Name    string               `json:"name"`
	Address jsonLDPostalAddress  `json:"address"`
	Geo     jsonLDGeoCoordinates `json:"geo"`
	URL     string               `json:"url"`
	Image   string               `json:"image,omitempty"`
	SameAs  []string             `json:"sameAs,omitempty"`
}

// placeJSONLD будує JSON-LD розмітку заведення (Restaurant/FoodEstablishment)
// для сторінки заклада: тільки реальні заповнені поля (адреса, геокоординати,
// соцмережі) — редакторські поля (design.md §4.6) не описуються, бо порожні.
// Для інших рівнів — порожньо.
func placeJSONLD(v any) template.JS {
	place, ok := v.(domain.Place)
	if !ok {
		return ""
	}

	var sameAs []string
	if place.InstagramURL != "" {
		sameAs = append(sameAs, place.InstagramURL)
	}
	if place.FacebookURL != "" {
		sameAs = append(sameAs, place.FacebookURL)
	}

	return marshalJSONLD(jsonLDPlace{
		Context: "https://schema.org",
		Type:    schemaOrgPlaceType(place.Type),
		Name:    place.Name,
		Address: jsonLDPostalAddress{
			Type: "PostalAddress",
			// StreetAddress у даних — латиницею (напр. "…Nicosia, Cyprus"),
			// тож AddressLocality/AddressCountry теж латиницею й кодом ISO,
			// щоб не змішувати мови в одній адресі (задача 3.3.3). Видима
			// кирилична назва міста лишається в HTML/title/description.
			StreetAddress:   place.Address,
			AddressLocality: cityLatinName(place.City.Alias),
			AddressCountry:  place.City.Country.ISOCode,
		},
		Geo: jsonLDGeoCoordinates{
			Type:      "GeoCoordinates",
			Latitude:  place.Location.Latitude,
			Longitude: place.Location.Longitude,
		},
		URL:    absURL(PlaceURL(place)),
		// Image — той самий og:image (ogImage вище): порожній PhotoURL дає
		// "", а omitempty прибирає поле з JSON-LD цілком.
		Image:  ogImage(place),
		SameAs: sameAs,
	})
}
