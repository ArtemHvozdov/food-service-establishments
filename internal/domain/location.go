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
	GermanyBerlin = City{
		Country: Germany,
		Alias:   "berlin",
		Name:    "Берлін",
	}
	GermanyBonn = City{
		Country: Germany,
		Alias:   "bonn",
		Name:    "Бонн",
	}
	GermanyBremen = City{
		Country: Germany,
		Alias:   "bremen",
		Name:    "Бремен",
	}
	GermanyDatteln = City{
		Country: Germany,
		Alias:   "datteln",
		Name:    "Даттельн",
	}
	GermanyDortmund = City{
		Country: Germany,
		Alias:   "dortmund",
		Name:    "Дортмунд",
	}
	GermanyDresden = City{
		Country: Germany,
		Alias:   "dresden",
		Name:    "Дрезден",
	}
	GermanyDusseldorf = City{
		Country: Germany,
		Alias:   "dusseldorf",
		Name:    "Дюссельдорф",
	}
	GermanyFrankfurtAmMain = City{
		Country: Germany,
		Alias:   "frankfurt-am-main",
		Name:    "Франкфурт-на-Майні",
	}
	GermanyHannover = City{
		Country: Germany,
		Alias:   "hannover",
		Name:    "Ганновер",
	}
	GermanyKarlsruhe = City{
		Country: Germany,
		Alias:   "karlsruhe",
		Name:    "Карлсруе",
	}
	GermanyKoln = City{
		Country: Germany,
		Alias:   "koln",
		Name:    "Кельн",
	}
	GermanyMunchen = City{
		Country: Germany,
		Alias:   "munchen",
		Name:    "Мюнхен",
	}
	GermanyNurnberg = City{
		Country: Germany,
		Alias:   "nurnberg",
		Name:    "Нюрнберг",
	}
	GermanyPullachImIsartal = City{
		Country: Germany,
		Alias:   "pullach-im-isartal",
		Name:    "Пуллах-Ізарталь",
	}
	GermanyRostock = City{
		Country: Germany,
		Alias:   "rostock",
		Name:    "Росток",
	}
	GermanyStrubbel = City{
		Country: Germany,
		Alias:   "strubbel",
		Name:    "Штрюббель",
	}
	GermanyStuttgart = City{
		Country: Germany,
		Alias:   "stuttgart",
		Name:    "Штутгарт",
	}
	SpainAlicante = City{
		Country: Spain,
		Alias:   "alicante",
		Name:    "Аліканте",
	}
	SpainBarcelona = City{
		Country: Spain,
		Alias:   "barcelona",
		Name:    "Барселона",
	}
	SpainBenalmadena = City{
		Country: Spain,
		Alias:   "benalmadena",
		Name:    "Бенальмадена",
	}
	SpainCurtis = City{
		Country: Spain,
		Alias:   "curtis",
		Name:    "Куртіс",
	}
	SpainMadrid = City{
		Country: Spain,
		Alias:   "madrid",
		Name:    "Мадрид",
	}
	SpainSantFeliuDeLlobregat = City{
		Country: Spain,
		Alias:   "sant-feliu-de-llobregat",
		Name:    "Сант-Феліу-де-Льобрегат",
	}
	SpainValencia = City{
		Country: Spain,
		Alias:   "valencia",
		Name:    "Валенсія",
	}
	HungaryBudapest = City{
		Country: Hungary,
		Alias:   "budapest",
		Name:    "Будапешт",
	}
	PolandGdansk = City{
		Country: Poland,
		Alias:   "gdansk",
		Name:    "Гданськ",
	}
	PolandGdynia = City{
		Country: Poland,
		Alias:   "gdynia",
		Name:    "Гдиня",
	}
	PolandKrakow = City{
		Country: Poland,
		Alias:   "krakow",
		Name:    "Краків",
	}
	PolandSzczecin = City{
		Country: Poland,
		Alias:   "szczecin",
		Name:    "Щецин",
	}
	PolandWarszawa = City{
		Country: Poland,
		Alias:   "warszawa",
		Name:    "Варшава",
	}
	PolandWroclaw = City{
		Country: Poland,
		Alias:   "wroclaw",
		Name:    "Вроцлав",
	}
	PortugalCascais = City{
		Country: Portugal,
		Alias:   "cascais",
		Name:    "Кашкайш",
	}
	PortugalCoimbra = City{
		Country: Portugal,
		Alias:   "coimbra",
		Name:    "Коїмбра",
	}
	PortugalLagos = City{
		Country: Portugal,
		Alias:   "lagos",
		Name:    "Лагуш",
	}
	PortugalLisboa = City{
		Country: Portugal,
		Alias:   "lisboa",
		Name:    "Лісабон",
	}
	PortugalPontinha = City{
		Country: Portugal,
		Alias:   "pontinha",
		Name:    "Понтінья",
	}
)
