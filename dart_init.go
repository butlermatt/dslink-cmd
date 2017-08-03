package main

import (
	"os"
	"fmt"
	"text/template"

	"github.com/pkg/errors"
	"github.com/butlermatt/dslink-cmd/templates"
	"path/filepath"
)

func initDart(c initConf) error {
	if c.Dir == "" {
		c.Dir = fmt.Sprintf("dslink-%s-%s", c.Lang, c.Name)
	}

	root := &FileItem{Path: c.Dir, IsDir: true}
	root.Add(&FileItem{Path: filepath.Join(c.Dir, "README.md"), Tmpl: templates.README})
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

func walkFiles(f *FileItem, c initConf) error {
	fmt.Println("Working on:", f.Path)
	fmt.Printf("Children: %+v", f.Childs)
	if !f.IsDir {
		return mkfile(f.Path, f.Tmpl, c)
	}

	err := os.Mkdir(f.Path, 0755)
	if err != nil {
		return errors.Wrapf(err, "error creating directory %q", f.Path)
	}

	for _, ff := range f.Childs {
		err := walkFiles(ff, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func mkfile(fn, tmpl string, conf initConf) (err error) {
	file, err := os.Create(fn)
	if err != nil {
		err = errors.Wrapf(err, "unable to generate file: %s", fn)
		return err
	}
	defer func() {
		e := file.Close()
		if e != nil {
			err = errors.Wrapf(e, "error closing file %s", file.Name())
		}
		fmt.Println("closed file: ", file.Name())
	}()

	tmplate, err := template.New(fn).Parse(tmpl)
	if err != nil {
		err = errors.Wrapf(err, "error parsing template for file %q", fn)
		return err
	}
	err = tmplate.Execute(file, conf)
	if err != nil {
		err = errors.Wrap(err, "error executing template")
		return err
	}

	return nil
}