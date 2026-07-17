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
	CyprusKouklia = City{
		Country: Cyprus,
		Alias:   "kouklia",
		Name:    "Куклія",
	}
	CyprusLarnaca = City{
		Country: Cyprus,
		Alias:   "larnaca",
		Name:    "Ларнака",
	}
	CyprusNicosia = City{
		Country: Cyprus,
		Alias:   "nicosia",
		Name:    "Нікосія",
	}
	CzechiaBrno = City{
		Country: Czechia,
		Alias:   "brno",
		Name:    "Брно",
	}
	CzechiaJihlava = City{
		Country: Czechia,
		Alias:   "jihlava",
		Name:    "Їглава",
	}
	CzechiaKarlovyVary = City{
		Country: Czechia,
		Alias:   "karlovy-vary",
		Name:    "Карлові Вари",
	}
	CzechiaLiberec = City{
		Country: Czechia,
		Alias:   "liberec",
		Name:    "Ліберец",
	}
	CzechiaMladaBoleslav = City{
		Country: Czechia,
		Alias:   "mlada-boleslav",
		Name:    "Млада-Болеслав",
	}
	CzechiaOlomouc = City{
		Country: Czechia,
		Alias:   "olomouc",
		Name:    "Оломоуц",
	}
	CzechiaPardubice = City{
		Country: Czechia,
		Alias:   "pardubice",
		Name:    "Пардубице",
	}
	CzechiaPraha = City{
		Country: Czechia,
		Alias:   "praha",
		Name:    "Прага",
	}
	CzechiaVranovNadDyji = City{
		Country: Czechia,
		Alias:   "vranov-nad-dyji",
		Name:    "Вранов-над-Дийї",
	}
	CzechiaUstiNadLabem = City{
		Country: Czechia,
		Alias:   "usti-nad-labem",
		Name:    "Усті-над-Лабем",
	}
)
