# gopackager

Command line tool that will package a project to allow serving as a gomodule (see below)

With the introduction of Go Modules the tooling is well positioned for proxying packages. 
The issue https://github.com/golang/go/issues/33312 is requesting for a way to zip up a project in the GoModule format allowing package managers
such as Nexus Repository Manager to locally push GoModule content.

## Program arguments
`-src` path location of the project you want to package (default `.`)

`-dst` path location to store the packaged project (default `.`)

`-project` defines the project name reference, e.g. github.com/doddi/gopackager

`-version` version of the application, e.g. `v1.0.0` 