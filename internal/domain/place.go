package domain

import "time"

type PlaceType struct {
	Alias string // restaurant
	Name  string // Ресторан
}

var (
	Restaurant = PlaceType{
		Alias: "restaurant",
		Name:  "Ресторан",
	}
	Cafe = PlaceType{
		Alias: "cafe",
		Name:  "Кафе",
	}
	Bistro = PlaceType{
		Alias: "bistro",
		Name:  "Бістро",
	}
	Bar = PlaceType{
		Alias: "bar",
		Name:  "Бар",
	}
	Bakery = PlaceType{
		Alias: "bakery",
		Name:  "Пекарня",
	}
)

type Place struct {
	Alias           string   // Domain name or Instagram username, e.g. "borshch.krakow"
	PreviousAliases []string // For redirect from old aliases to current one
	City            City
	Name            string
	Address         string
	Type            PlaceType
	URL             string
	MenuURL         string
	PhotoURL        string
	InstagramURL    string
	FacebookURL     string
	GoogleMapsURL   string
	GooglePlaceID   string
	Location        Location
	FoundingYear    int
	Description     string
	Dishes          []Dish
	BorschPrice     string // value and currency, e.g. "50 UAH"
	LargePortions   bool
	HasDelivery     bool
	DonatesToZSU    []Article
	Articles        []Article
	YoutubeVideos   []YoutubeVideo
	AddedAt         time.Time
	PromotedUntil   *time.Time
	// TODO: add ratings TripAdvisor, Google, etc.
}

type CityPlaceGroup struct {
	City   City
	Places []Place
}

type CountryPlaceGroup struct {
	Country Country
	Cities  []CityPlaceGroup
}
