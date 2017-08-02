package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

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

var (
	app = kingpin.New(os.Args[0], "A command line project maintainer for DSLinks")

	build       = app.Command("build", "Build an existing project.")
	buildType   = build.Arg("language", "Programming language of the project to build.").Enum(buildTypes...)
	buildOutput = build.Flag("out", "Output file").Short('o').Default("build.zip").String()
)

func main() {
	var err error

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case build.FullCommand():
		if *buildType == "" {
			*buildType, err = checkBuild()
		}
		if err == nil {
			fmt.Println("build: ", *buildType)
			err = tryBuild(*buildType, *buildOutput)
		}
	}

	if err != nil {
		fmt.Println("error occurred:", err)
	} else {
		fmt.Println("Done!")
	}
}

func tryBuild(ty, out string) error {
	var err error
	switch ty {
	case bDart:
		err = buildDart(out)
	case bC:
	case bDotNet:
	case bJava:
	case bJS:
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

func checkBuild() (string, error) {
	var cantFind error = fmt.Errorf("unable to determine project type. Please specify build language.")

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cd := filepath.Base(dir)
	fmt.Println(cd)

	si := strings.Index(cd, "-")
	if si == -1 {
		return "", cantFind
	}
	if len(cd) <= si + 1 {
		return "", cantFind
	}

	ei := strings.Index(cd[si + 1:], "-")
	if ei == -1 {
		return "", cantFind
	}

	return cd[si + 1:ei + si + 1], nil
}
