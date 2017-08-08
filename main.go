package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "dslink-cmd version: 0.0.2"

const (
	bDart   string = "dart"
	bC             = "c"
	bDotNet        = "dotnet"
	bJava          = "java"
	bJS            = "javascript"
	bPython        = "python"
	bScala         = "scala"
)

var buildTypes = []string{bDart, bJava, bJS, bPython, bDotNet, bC, bScala}

type initConf struct {
	Lang string
	Name string
	Node string
	Dir  string
}

var (
	app = kingpin.New(os.Args[0], "A command line project maintainer for DSLinks").Version(version)

	buildCmd    = app.Command("build", "Build an existing project.")
	buildType   = buildCmd.Arg("language", "Programming language of the project to build.").Enum(buildTypes...)
	buildOutput = buildCmd.Flag("out", "Output file").Short('o').Default("build.zip").String()

	initCmd  = app.Command("init", "Initialize a new project")
	initType = initCmd.Arg("language", "Programming language of the project to initialize.").Required().Enum(buildTypes...)
	initName = initCmd.Arg("name", "Short name of the project.").Required().String()
	initNode = initCmd.Flag("node", "Default node name in the DSA hierarchy. Uses program name if unspecified.").Short('n').String()
	initDir  = initCmd.Flag("dir", "Directory to initialize the project into. Defaults to dslink-<langauge>-<name>").Short('d').String()
)

func main() {
	var err error

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case buildCmd.FullCommand():
		if *buildType == "" {
			*buildType, err = checkCWD()
		}
		if err == nil {
			err = tryBuild(*buildType, *buildOutput)
		}
	case initCmd.FullCommand():
		if *initNode == "" {
			*initNode = *initName
		}
		err = tryInit(initConf{*initType, *initName, *initNode, *initDir})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func tryBuild(ty, out string) error {
	var err error

	switch ty {
	case bDart:
		err = buildDart(out)
	case bJS:
	case bC:
	case bDotNet:
	case bJava:
	case bPython:
	case bScala:
		err = fmt.Errorf("build for %q is not currently implemented.\n", ty)
	default:
		err = fmt.Errorf("unknown build language %q", ty)
	}

	if err != nil {
		err = errors.Wrap(err, "error building")
	}

	return err
}

func checkCWD() (string, error) {
	var cantFind error = fmt.Errorf("unable to determine project type. Please specify build language.")

	dir, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "error checking working directory")
	}

	cd := filepath.Base(dir)

	si := strings.Index(cd, "-")
	if si == -1 {
		return "", cantFind
	}

	si += 1 // bump to just after start index
	if len(cd) <= si {
		return "", cantFind
	}

	ei := strings.Index(cd[si:], "-")
	if ei == -1 {
		return "", cantFind
	}

	return cd[si : ei+si], nil
}

func tryInit(conf initConf) error {
	var err error

	switch conf.Lang {
	case bDart:
		err = initDart(conf)
	case bJS:
		err = initJs(conf)
	case bC:
	case bDotNet:
	case bJava:
	case bPython:
	case bScala:
		err = fmt.Errorf("init for %q is not currently implemented.\n", conf.Lang)
	default:
		err = fmt.Errorf("unknown init language %q", conf.Lang)
	}

	if err != nil {
		err = errors.Wrap(err, "error initializing")
	}

	return err
}
