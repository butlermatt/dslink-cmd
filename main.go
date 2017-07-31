package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
)

var (
	buildTypes = []string{"dart"}
)

var (
	app = kingpin.New(os.Args[0], "A command line project maintainer for DSLinks")

	build = app.Command("build", "Build an existing project.")
	buildType = build.Arg("language", "Programming language of the project to build.").Required().Enum(buildTypes...)
	buildOutput = build.Flag("out", "Output file").Short('o').Default("build.zip").String()
)

func main() {
	var err error
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case build.FullCommand():
		switch *buildType {
		case buildTypes[0]:
			err = buildDart(*buildOutput)
		}
	}

	if err != nil {
		fmt.Println("error occurred:", err)
	} else {
		fmt.Println("Done!")
	}
}