package packager

import (
	"fmt"
	"github.com/doddi/gopackager/packager/compress"
	"github.com/doddi/gopackager/packager/gomodule"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func contains(path string, files ...string) bool {

	found := 0

	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// already found what we need, skip
		if found >= len(files) {
			return nil
		}

		for _, file := range files {
			if path == file {
				found++
				if found >= len(files) {
					break
				}
			}
		}
		return nil
	})

	if found >= len(files) {
		return true
	}
	return false
}

func Package(goModule gomodule.GoModule, sourcePath string, destinationPath string) {
	validateProject(sourcePath)

	//TODO If the directories already exist, fail

	defer removeTempFolder(goModule.GetProjectPath())

	// Create temp folder to contain temporary copy of project
	fullPathToTempProject := createTemporaryFolder(goModule.GetProjectPath())

	copyProject(sourcePath, fullPathToTempProject)
	compress.Folder(
		os.TempDir()+goModule.GetVcs(),
		destinationPath+string(os.PathSeparator)+goModule.GetModuleZipName(),
	)
}

func validateProject(sourcePath string) {
	if !contains(sourcePath, "go.mod") {
		fmt.Println("Failed to find project's go.mod file")
		os.Exit(1)
	}
}

func createTemporaryFolder(path string) string {
	absoluteDestination := os.TempDir() + path
	err := os.MkdirAll(absoluteDestination, os.ModeDir|os.ModePerm)
	if err != nil {
		panic(err)
	}
	return absoluteDestination
}

func copyProject(sourcePath string, destinationPath string) {
	filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		dst := destinationPath + string(os.PathSeparator) + path
		src := sourcePath + string(os.PathSeparator) + path
		if info.IsDir() {
			os.Mkdir(dst, os.ModeDir|os.ModePerm)
			return nil
		}
		copyFile(src, dst)
		return nil
	})
}

func copyFile(source string, destination string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func removeTempFolder(destinationPath string) {
	fmt.Println("Cleaning up temporary folders")

	split := strings.Split(destinationPath, string(os.PathSeparator))

	err := os.RemoveAll(os.TempDir() + split[0])
	if err != nil {
		panic(err)
	}
}
