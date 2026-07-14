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

// Реестр країн (код -> Country), див. CLAUDE.md та internal_docs/task_01.md,
// Додаток A. Країни визначаються вручну: множина маленька і майже не росте.
var (
	Bulgaria = Country{
		Alias: "bulgaria",
		Name:  "Болгарія",
	}
	Cyprus = Country{
		Alias: "cyprus",
		Name:  "Кіпр",
	}
	Czechia = Country{
		Alias: "czechia",
		Name:  "Чехія",
	}
	Germany = Country{
		Alias: "germany",
		Name:  "Німеччина",
	}
	Spain = Country{
		Alias: "spain",
		Name:  "Іспанія",
	}
	Poland = Country{
		Alias: "poland",
		Name:  "Польща",
	}
	Portugal = Country{
		Alias: "portugal",
		Name:  "Португалія",
	}
	Romania = Country{
		Alias: "romania",
		Name:  "Румунія",
	}
	Slovakia = Country{
		Alias: "slovakia",
		Name:  "Словаччина",
	}
	Turkiye = Country{
		Alias: "turkiye",
		Name:  "Туреччина",
	}
	Hungary = Country{
		Alias: "hungary",
		Name:  "Угорщина",
	}
)

var (
	BulgariaVarna = City{
		Country: Bulgaria,
		Alias:   "varna",
		Name:    "Варна",
	}
	BulgariaSofia = City{
		Country: Bulgaria,
		Alias:   "sofia",
		Name:    "Софія",
	}
)
