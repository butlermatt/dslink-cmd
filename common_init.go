package main

import (
	"os"
	"fmt"
	"text/template"
	"path/filepath"

	"github.com/pkg/errors"
)

type FileItem struct {
	Path   string
	Tmpl   string
	IsDir  bool
	Childs []*FileItem
	parent *FileItem
}

func (f *FileItem) Add(fi *FileItem) {
	fi.parent = f
	f.Childs = append(f.Childs, fi)
}

func (f *FileItem) FilePath() string {
	var p string
	if f.parent != nil {
		p = f.parent.FilePath()
	}

	return filepath.Join(p, f.Path)
}

func walkFiles(f *FileItem, c initConf) error {
	if !f.IsDir {
		return mkfile(f.FilePath(), f.Tmpl, c)
	}

	err := os.Mkdir(f.FilePath(), 0755)
	if err != nil {
		return errors.Wrapf(err, "error creating directory %q", f.FilePath())
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