// Package assets вшиває сирі JSON-дані (data/*.json) у бінарник через
// go:embed, щоб завантаження не залежало від робочої директорії запуску
// (internal_docs/task_01.md, Підзадача 1.2). Директива embed не дозволяє
// ".." у шляху, тому файл лежить у корені модуля — найближчому спільному
// предку для data/ та internal/db.
package assets

import "embed"

//go:embed data/places/*.json data/countries.json
var FS embed.FS
