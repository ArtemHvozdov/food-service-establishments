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

// photosFS — фото закладів (static/photos/<country>/<city>/<place>.webp),
// вбудовані окремо від staticFS: go:embed не підтримує порожній патерн
// (ламає збірку), тому embed додається лише коли папка непорожня.
//
//go:embed static/photos
var photosFS embed.FS

// copyStaticAssets копіює вміст staticFS і photosFS у корінь outputDir,
// зберігаючи відносні шляхи:
//   static/style.css             → outputDir/style.css
//   static/fonts/*.woff2         → outputDir/fonts/*.woff2
//   static/photos/**/*.webp      → outputDir/photos/**/*.webp
func copyStaticAssets(outputDir string) error {
	if err := copyFS(staticFS, "static", outputDir); err != nil {
		return err
	}
	return copyFS(photosFS, "static/photos", filepath.Join(outputDir, "photos"))
}

// copyFS обходить fsys починаючи з root і копіює всі файли у destDir,
// зберігаючи відносну структуру підпапок.
func copyFS(fsys embed.FS, root, destDir string) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		dst := filepath.Join(destDir, rel)
		data, err := fsys.ReadFile(path)
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
