package main

import (
	"fmt"

	"github.com/butlermatt/dslink-cmd/templates"
)

func initDart(c initConf) error {
	if c.Dir == "" {
		c.Dir = fmt.Sprintf("dslink-%s-%s", c.Lang, c.Name)
	}

	root := &FileItem{Path: c.Dir, IsDir: true}
	root.Add(&FileItem{Path: "README.md", Tmpl: templates.README})
	root.Add(&FileItem{Path: "LICENSE", Tmpl: templates.LICENSE})
	root.Add(&FileItem{Path: ".gitignore", Tmpl: templates.GitIgnore})
	root.Add(&FileItem{Path: "dslink.json", Tmpl: templates.DSLinkJson})
	root.Add(&FileItem{Path: "pubspec.yaml", Tmpl: templates.DartPubSpec})

	bin := &FileItem{Path: "bin", IsDir: true}
	bin.Add(&FileItem{Path: "run.dart", Tmpl: templates.DartRun})
	root.Add(bin)

	lib := &FileItem{Path: "lib", IsDir: true}
	lib.Add(&FileItem{Path: c.Name + ".dart"})
	lib.Add(&FileItem{Path: "models.dart"})
	root.Add(lib)

	src := &FileItem{Path: "src", IsDir: true}
	src.Add(&FileItem{Path: "models", IsDir: true})
	src.Add(&FileItem{Path: "nodes", IsDir: true})
	lib.Add(src)

	return walkFiles(root, c)
}
