package db

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// slugOverrides — випадки, де авто-згортка через NFD дає невірний результат
// (насамперед кирилиця, яка не розкладається на базову літеру + діакритику).
// Перевіряється першою, до решти алгоритму.
var slugOverrides = map[string]string{
	"Варна": "varna",
	"София": "sofia",
}

// slugRuneReplacements — літери, які NFD не розкладає на базову літеру +
// комбінуючий знак (це не діакритика, а окрема буква з власним накресленням).
var slugRuneReplacements = map[rune]string{
	'ł': "l", 'Ł': "l",
	'ß': "ss", 'ẞ': "ss",
	'ø': "o", 'Ø': "o",
	'đ': "d", 'Đ': "d",
}

// cyrillicTransliteration — транслітерація кирилиці в латиницю. Ключі —
// лише строчні руни; для великих літер робимо lookup через unicode.ToLower,
// щоб не дублювати записи. Точність не лінгвістична — важливі детермінізм
// і непорожній ASCII-результат (internal_docs/task_01.md, 1.3.1).
//
// й/ї/ё в NFD канонічно розкладаються на «базова літера + комбінуючий
// діакритичний знак» (й → и+breve, ї → і+diaeresis, ё → е+diaeresis), тож до
// цієї мапи вони доходять вже розкладеними на свою базову літеру (Mn-знак
// прибирається раніше в циклі). Записи для й/ї/ё тут — підстраховка на
// випадок NFC/нерозкладеного вводу, і мають збігатися з базовою літерою:
// й→y (як и→y), ї→i (як і→i), ё→e (як е→e).
var cyrillicTransliteration = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "h", 'ґ': "g", 'д': "d", 'е': "e",
	'є': "ie", 'ж': "zh", 'з': "z", 'и': "y", 'і': "i", 'ї': "i", 'й': "y",
	'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o", 'п': "p", 'р': "r",
	'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch",
	'ш': "sh", 'щ': "shch", 'ъ': "", 'ь': "", 'ю': "iu", 'я': "ia", 'ы': "y",
	'э': "e", 'ё': "e",
}

// slugify перетворює нативну назву (міста, закладу тощо) на URL-аліас:
// lowercase, ASCII [a-z0-9-], без крайніх дефісів. Єдиний алгоритм для
// міст і закладів, щоб URL-схема сайту була передбачуваною.
func slugify(s string) string {
	if override, ok := slugOverrides[s]; ok {
		return override
	}

	normalized := norm.NFD.String(s)

	var stripped strings.Builder
	for _, r := range normalized {
		if unicode.Is(unicode.Mn, r) {
			continue // прибираємо комбінуючі діакритичні знаки
		}
		if repl, ok := cyrillicTransliteration[unicode.ToLower(r)]; ok {
			stripped.WriteString(repl)
			continue
		}
		if repl, ok := slugRuneReplacements[r]; ok {
			stripped.WriteString(repl)
			continue
		}
		stripped.WriteRune(r)
	}

	lowered := strings.ToLower(stripped.String())

	var out strings.Builder
	prevDash := false
	for _, r := range lowered {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			out.WriteRune(r)
			prevDash = false
		case !prevDash:
			out.WriteByte('-')
			prevDash = true
		}
	}

	return strings.Trim(out.String(), "-")
}
