package main

import (
	"flag"
	"github.com/doddi/gopackager/packager"
	"github.com/doddi/gopackager/packager/gomodule"
)

func main() {
	var source, destination, project, version string

	flag.StringVar(&source, "src", ".", "source path of go project. default \".\"")
	flag.StringVar(&destination, "dst", ".", "destination path of go zip, default \".\"")
	flag.StringVar(&version, "version", "", "version of your go project")
	flag.StringVar(&project, "project", "", "project name <vcs>/<owner>/<project>")
	flag.Parse()

	goModule, err := gomodule.Parse(project, version)
	if err != nil {
		panic("Unable to parse go module name provided")
	}

	packager.Package(goModule, source, destination)
}
