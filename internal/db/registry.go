package db

import "github.com/ArtemHvozdov/food-service-establishments/internal/domain"

// countryByCode — резолвинг коду країни з JSON (напр. "pl") у domain.Country.
// Довідник статичний: набір країн маленький і майже не росте (internal_docs/task_01.md, Додаток A).
var countryByCode = map[string]domain.Country{
	"bg": domain.Bulgaria,
	"cy": domain.Cyprus,
	"cz": domain.Czechia,
	"de": domain.Germany,
	"es": domain.Spain,
	"pl": domain.Poland,
	"pt": domain.Portugal,
	"ro": domain.Romania,
	"sk": domain.Slovakia,
	"tr": domain.Turkiye,
	"hu": domain.Hungary,
}

// placeTypeByString — резолвинг рядка type з JSON у domain.PlaceType.
var placeTypeByString = map[string]domain.PlaceType{
	"restaurant": domain.Restaurant,
	"cafe":       domain.Cafe,
	"bistro":     domain.Bistro,
	"bar":        domain.Bar,
	"bakery":     domain.Bakery,
}
