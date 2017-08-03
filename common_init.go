package main

import (
	"os"
	"fmt"
	"text/template"

	"github.com/pkg/errors"
)

type FileItem struct {
	Path   string
	Tmpl   string
	IsDir  bool
	Childs []*FileItem
}

func (f *FileItem) Add(fi *FileItem) {
	f.Childs = append(f.Childs, fi)
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