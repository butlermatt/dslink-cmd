package main

import (
	"fmt"
	"path/filepath"

	"github.com/butlermatt/dslink-cmd/templates"
)

func initJs(c initConf) error {
	if c.Dir == "" {
		c.Dir = fmt.Sprintf("dslink-%s-%s", c.Lang, c.Name)
	}

	root := &FileItem{Path: c.Dir, IsDir: true}
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "README.md"), Tmpl: templates.README})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "LICENSE"), Tmpl: templates.LICENSE})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "dslink.json"), Tmpl: templates.DSLinkJson})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, ".gitignore"), Tmpl: templates.GitIgnore})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "package.json"), Tmpl: templates.JsPackageJson})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "install.js"), Tmpl: templates.JsInstall})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "index.js"), Tmpl: templates.JsIndex})

	return walkFiles(root, c)
}
