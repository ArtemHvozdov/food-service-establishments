package domain

type Location struct {
	Latitude  float64
	Longitude float64
}

type Country struct {
	Alias string // Poland
	Name  string // Польща
}

type City struct {
	Country Country
	Alias   string // krakow
	Name    string // Краків
}

var (
	Poland = Country{
		Alias: "poland",
		Name:  "Польща",
	}
	// ...

	Krakow = City{
		Country: Poland,
		Alias:   "krakow",
		Name:    "Краків",
	}
)
