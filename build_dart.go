package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"bytes"
)

const (
	pkgsFile = ".packages"
	pkgDir   = "pkg"
)

var ignoreList = []string{".git", ".gitignore", ".gitmodules", ".arc", ".idea"}

func buildDart(file string) error {
	ignoreList = append(ignoreList, file)

	z, err := os.Create(file)
	if err != nil {
		return err
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	err = filepath.Walk(".", walker("","", w))

	return err
}

func cp(src, dest string, out *zip.Writer) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	fout, err := out.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(fout, in)
	if err != nil {
		return err
	}

	return nil
}

func processPackages(file string, out *zip.Writer) error {
	pf, err := os.Open(file)
	if err != nil {
		return err
	}
	defer pf.Close()

	w := new(bytes.Buffer)

	//w := bufio.NewWriter(outf)

	scan := bufio.NewScanner(pf)
	for scan.Scan() {
		ln := scan.Text()
		if ln[0] == '#' {
			continue
		}

		p := strings.SplitN(ln, ":", 2)
		name := p[0]
		addr, err := url.Parse(p[1])
		if err != nil {
			return fmt.Errorf("error parsing .package url: %v", err)
		}

		// This should only be the case if it's already relative
		if addr.Scheme == "" {
			_, err = w.WriteString(ln)
			if err != nil {
				return fmt.Errorf("error writing to zip .package: %v", err)
			}

			continue
		}

		pth := filepath.Dir(filepath.Dir(addr.Path))
		err = filepath.Walk(pth, walker(pth, pkgDir, out))
		if err != nil {
			return err
		}

		root := filepath.Base(pth)
		ind := strings.Index(addr.Path, root)
		fp := filepath.Join(pkgDir, addr.Path[ind:])
		if fp[len(fp) - 1] != '/' {
			fp = fp + "/"
		}
		_, err = w.WriteString(strings.Join([]string{name, fp}, ":") + "\n")
		if err != nil {
			return err
		}
	}

	outf, err := out.Create(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(outf, w)
	if err != nil {
		return err
	}

	return nil
}

func walker(root, prefix string, out *zip.Writer) (func(string, os.FileInfo, error) error) {
	var top string
	if prefix != "" {
		top = filepath.Base(root)
	}

	return func(path string, info os.FileInfo, err error) error {
		for _, ignore := range ignoreList {
			if info.Name() == ignore {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		// Prefix is an easy check to make sure we're not looking at a .package file in the package cache.
		if info.Name() == pkgsFile && prefix == "" {
			err = processPackages(path, out)
			if err != nil {
				return err
			}
		} else {
			dst := path
			if prefix != "" {
				ind := strings.Index(path, top)
				dst = filepath.Join(prefix, path[ind:])
			}
			cp(path, dst, out)
		}

		if prefix == "" {
			fmt.Printf("file: %s\n", path)
		}
		return nil
	}
}
