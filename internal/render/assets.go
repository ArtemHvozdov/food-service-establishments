package render

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// staticFS — статичні ассети (CSS, self-hosted шрифт), що доставляються у
// корінь outputDir разом зі сторінками (design.md §6): public/ лишається
// артефактом Generate(), ассети в нього не кладуться руками.
//
//go:embed static/style.css static/fonts/*.woff2
var staticFS embed.FS

// copyStaticAssets копіює вміст staticFS у корінь outputDir, зберігаючи
// відносні шляхи (static/style.css → outputDir/style.css,
// static/fonts/*.woff2 → outputDir/fonts/*.woff2) — у розмітці ці шляхи
// кореневі (/style.css, /fonts/...).
func copyStaticAssets(outputDir string) error {
	return fs.WalkDir(staticFS, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel("static", path)
		if err != nil {
			return err
		}

		dst := filepath.Join(outputDir, rel)
		data, err := staticFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("render: read asset %s: %w", path, err)
		}

		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return fmt.Errorf("render: mkdir %s: %w", filepath.Dir(dst), err)
		}
		if err := os.WriteFile(dst, data, 0o644); err != nil {
			return fmt.Errorf("render: write asset %s: %w", dst, err)
		}

		return nil
	})
}
