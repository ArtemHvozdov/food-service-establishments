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
