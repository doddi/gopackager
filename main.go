package main

import (
	"fmt"
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
	projectName := "github.com/doddi/gopackager"
	projectVersion := "1.0.0"
	destinationPath := "."
	packager.Package(sourcePath, projectName, projectVersion, destinationPath)
}
