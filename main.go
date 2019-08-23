package main

import (
	"fmt"
	"gopackager/gomodule"
	"gopackager/packager"
	"os"
)

func main() {

	commandLineArguments := os.Args[1:]

	if len(commandLineArguments) == 0 {
		fmt.Println("Please provide:\n\tProject name in the form  <vcs>/<projectname> e.g. github.com/sonatype/example\n\tVersion\n\tDirectory location to compress")
		os.Exit(1)
	}

	sourcePath := "."
	//projectName := "github.com/doddi/gopackager"
	//projectVersion := "1.0.0"
	destinationPath := "."

	goModule, err := gomodule.Parse("github.com/doddi/gopackager", "v0.0.1")
	if err != nil {
		panic("Unable to parse go module name provided")
	}

	packager.Package(sourcePath, goModule, destinationPath)
}
