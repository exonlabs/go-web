package webui

import (
	"go/build"
	"os"
	"path/filepath"
)

func PathsSetup(static string) error {
	staticWebUi := filepath.Join(static, "webui")
	tmplWebUi := filepath.Join("templates", "webui")
	pkgName := "github.com/exonlabs/go-web/pkg/webui"

	pkg, err := build.Default.Import(pkgName, build.Default.GOPATH,
		build.FindOnly)
	if err != nil {
		return err
	}

	// static operation
	if !exists(staticWebUi) {
		// create static webui dir
		if err := os.Mkdir(staticWebUi, os.ModePerm); err != nil {
			return err
		}
		// copy static from lib to app
		if err := copyDirectory(filepath.Join(pkg.Dir, "static"),
			staticWebUi); err != nil {
			return err
		}
	}

	// tmpl operation
	if !exists(tmplWebUi) {
		// create tmpl webui dir
		if err := os.Mkdir(tmplWebUi, os.ModePerm); err != nil {
			return err
		}
		// copy tmpls from lib to app
		if err := copyDirectory(filepath.Join(pkg.Dir, "templates"),
			tmplWebUi); err != nil {
			return err
		}
	}

	return nil
}
