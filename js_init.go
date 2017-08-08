package main

import (
	"fmt"

	"github.com/butlermatt/dslink-cmd/templates"
)

func initJs(c initConf) error {
	if c.Dir == "" {
		c.Dir = fmt.Sprintf("dslink-%s-%s", c.Lang, c.Name)
	}

	root := &FileItem{Path: c.Dir, IsDir: true}
	root.Add(&FileItem{Path: "README.md", Tmpl: templates.README})
	root.Add(&FileItem{Path: "LICENSE", Tmpl: templates.LICENSE})
	root.Add(&FileItem{Path: "dslink.json", Tmpl: templates.DSLinkJson})
	root.Add(&FileItem{Path: ".gitignore", Tmpl: templates.GitIgnore})
	root.Add(&FileItem{Path: "package.json", Tmpl: templates.JsPackageJson})
	root.Add(&FileItem{Path: "install.js", Tmpl: templates.JsInstall})
	root.Add(&FileItem{Path: "index.js", Tmpl: templates.JsIndex})

	return walkFiles(root, c)
}
