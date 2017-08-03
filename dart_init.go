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

	err := os.Mkdir(c.Dir, 0755)
	if err != nil {
		if os.IsExist(err) {
			return errors.New("can't initialize on existing directory")
		} else {
			return errors.Wrapf(err, "error creating directory %q", c.Dir)
		}
	}

	err = mkfile(filepath.Join(c.Dir, "README.md"), templates.README, c)
	if err != nil {
		return err
	}
	err = mkfile(filepath.Join(c.Dir, "dslink.json"), templates.DSLinkJson, c)
	if err != nil {
		return err
	}
	err = mkfile(filepath.Join(c.Dir, "pubspec.yaml"), templates.DartPubSpec, c)
	if err != nil {
		return err
	}

	bin := filepath.Join(c.Dir, "bin")
	err = os.Mkdir(bin, 0755)
	if err != nil {
		return errors.Wrapf(err, "error creating directory %q", bin)
	}

	err = mkfile(filepath.Join(bin, "run.dart"), templates.DartRun, c)
	if err != nil {
		return err
	}

	lib := filepath.Join(c.Dir, "lib")
	err = os.Mkdir(lib, 0755)
	if err != nil {
		return errors.Wrapf(err, "error creating directory %q", lib)
	}

	err = mkfile(filepath.Join(lib, c.Name + ".dart"), "", c)
	if err != nil {
		return err
	}
	err = mkfile(filepath.Join(lib, "models.dart"), "", c)
	if err != nil {
		return err
	}

	models := filepath.Join(lib, "src", "models")
	err = os.MkdirAll(models, 0755)
	if err != nil {
		return errors.Wrapf(err, "error create directory %q", models)
	}

	nodes := filepath.Join(lib, "src", "nodes")
	err = os.MkdirAll(nodes, 0755)
	if err != nil {
		return errors.Wrapf(err, "error create directory %q", nodes)
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