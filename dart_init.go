package main

import (
	"fmt"
	"path/filepath"

	"github.com/butlermatt/dslink-cmd/templates"
)

func initDart(c initConf) error {
	if c.Dir == "" {
		c.Dir = fmt.Sprintf("dslink-%s-%s", c.Lang, c.Name)
	}

	root := &FileItem{Path: c.Dir, IsDir: true}
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "README.md"), Tmpl: templates.README})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "LICENSE"), Tmpl: templates.LICENSE})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, ".gitignore"), Tmpl: templates.GitIgnore})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "dslink.json"), Tmpl: templates.DSLinkJson})
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "pubspec.yaml"), Tmpl: templates.DartPubSpec})

	bin := &FileItem{Path: filepath.Join(c.Dir, "bin"), IsDir: true}
	bin.Add(&FileItem{Path: filepath.Join(bin.Path, "run.dart"), Tmpl: templates.DartRun})
	root.Add(bin)

	lib := &FileItem{Path: filepath.Join(c.Dir, "lib"), IsDir: true}
	lib.Add(&FileItem{Path: filepath.Join(lib.Path, c.Name + ".dart")})
	lib.Add(&FileItem{Path: filepath.Join(lib.Path, "models.dart")})
	root.Add(lib)

	src := &FileItem{Path: filepath.Join(lib.Path, "src"), IsDir: true}
	src.Add(&FileItem{Path: filepath.Join(src.Path, "models"), IsDir: true})
	src.Add(&FileItem{Path: filepath.Join(src.Path, "nodes"), IsDir: true})
	lib.Add(src)

	return walkFiles(root, c)
}
