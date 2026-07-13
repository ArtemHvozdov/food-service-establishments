package db

import "testing"

// Golden-тест для slugify: internal_docs/task_01.md, Додаток B (57 міст).
func TestSlugifyGolden(t *testing.T) {
	cases := []struct {
		country string
		city    string
		want    string
	}{
		{"bg", "Варна", "varna"},
		{"bg", "София", "sofia"},
		{"cy", "Kouklia", "kouklia"},
		{"cy", "Larnaca", "larnaca"},
		{"cy", "Nicosia", "nicosia"},
		{"cz", "Brno", "brno"},
		{"cz", "Jihlava", "jihlava"},
		{"cz", "Karlovy Vary", "karlovy-vary"},
		{"cz", "Liberec", "liberec"},
		{"cz", "Mladá Boleslav", "mlada-boleslav"},
		{"cz", "Olomouc", "olomouc"},
		{"cz", "Pardubice", "pardubice"},
		{"cz", "Praha", "praha"},
		{"cz", "Vranov nad Dyjí", "vranov-nad-dyji"},
		{"cz", "Ústí nad Labem", "usti-nad-labem"},
		{"de", "Berlin", "berlin"},
		{"de", "Bonn", "bonn"},
		{"de", "Bremen", "bremen"},
		{"de", "Datteln", "datteln"},
		{"de", "Dortmund", "dortmund"},
		{"de", "Dresden", "dresden"},
		{"de", "Düsseldorf", "dusseldorf"},
		{"de", "Frankfurt am Main", "frankfurt-am-main"},
		{"de", "Hannover", "hannover"},
		{"de", "Karlsruhe", "karlsruhe"},
		{"de", "Köln", "koln"},
		{"de", "München", "munchen"},
		{"de", "Nürnberg", "nurnberg"},
		{"de", "Pullach im Isartal", "pullach-im-isartal"},
		{"de", "Rostock", "rostock"},
		{"de", "Strübbel", "strubbel"},
		{"de", "Stuttgart", "stuttgart"},
		{"es", "Alicante", "alicante"},
		{"es", "Barcelona", "barcelona"},
		{"es", "Benalmádena", "benalmadena"},
		{"es", "Curtis", "curtis"},
		{"es", "Madrid", "madrid"},
		{"es", "Sant Feliu de Llobregat", "sant-feliu-de-llobregat"},
		{"es", "Valencia", "valencia"},
		{"hu", "Budapest", "budapest"},
		{"pl", "Gdańsk", "gdansk"},
		{"pl", "Gdynia", "gdynia"},
		{"pl", "Kraków", "krakow"},
		{"pl", "Szczecin", "szczecin"},
		{"pl", "Warszawa", "warszawa"},
		{"pl", "Wrocław", "wroclaw"},
		{"pt", "Cascais", "cascais"},
		{"pt", "Coimbra", "coimbra"},
		{"pt", "Lagos", "lagos"},
		{"pt", "Lisboa", "lisboa"},
		{"pt", "Pontinha", "pontinha"},
		{"ro", "București", "bucuresti"},
		{"ro", "Cluj-Napoca", "cluj-napoca"},
		{"ro", "Năvodari", "navodari"},
		{"sk", "Bratislava", "bratislava"},
		{"sk", "Košice", "kosice"},
		{"tr", "İzmir", "izmir"},
	}

	if got, want := len(cases), 57; got != want {
		t.Fatalf("golden test must cover 57 cities, got %d", got)
	}

	for _, tc := range cases {
		if got := slugify(tc.city); got != tc.want {
			t.Errorf("%s | slugify(%q) = %q, want %q", tc.country, tc.city, got, tc.want)
		}
	}
}

// TestSlugifyCyrillicTransliteration — регрес на кириличні назви закладів,
// які раніше давали порожній аліас (internal_docs/task_01.md, 1.3.1).
func TestSlugifyCyrillicTransliteration(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"Ресторан ТРОЯНДА", "restoran-troianda"},
		{"Файна Украина", "fayna-ukrayna"},
		{"Бистро Портокал", "bystro-portokal"},
	}

	for _, tc := range cases {
		if got := slugify(tc.name); got != tc.want {
			t.Errorf("slugify(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}
