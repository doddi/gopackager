package packager

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func contains(path string, files ...string) bool {

	found := 0

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
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

func Package(sourcePath string, projectName string, projectVersion string, destinationPath string) {
	if !contains(sourcePath, "go.mod") {
		fmt.Println("Not all files are present")
		os.Exit(1)
	}

	//TODO If the directories already exist, fail

	project := projectName + "@" + projectVersion

	defer removeTempFolder(project)

	// Create temp folder to contain project to copy
	fullPathToTempProject := createTemporaryFolder(project)

	copyProject(sourcePath, fullPathToTempProject)

	split := strings.Split(projectName, string(os.PathSeparator))
	projectZipName := split[len(split)-1] + ".zip"
	folderToZip := os.TempDir() + split[0]
	compressProject(project, folderToZip, projectZipName)
}

func compressProject(projectName string, folderToZip string, projectZipName string) {
	newZipFile, err := os.Create(projectZipName)
	if err != nil {
		panic(err)
	}

	writer := zip.NewWriter(newZipFile)
	defer writer.Close()

	filepath.Walk(folderToZip, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			addFileToZip(writer, path)
		}
		return nil
	})
}

func addFileToZip(writer *zip.Writer, file string) error {
	openedFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("unable to open file %v: %v", file, err)
	}
	defer openedFile.Close()

	// Strip the temp folder names
	fileToCreate := strings.Replace(file, os.TempDir(), "", 1)
	wr, err := writer.Create(fileToCreate)
	if err != nil {
		return fmt.Errorf("error adding file; '%s' to zip : %s", file, err)
	}

	if _, err := io.Copy(wr, openedFile); err != nil {
		return fmt.Errorf("error writing %s to zip: %s", file, err)
	}
	return nil
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
		dest := destinationPath + string(os.PathSeparator) + path
		if info.IsDir() {
			os.Mkdir(dest, os.ModeDir|os.ModePerm)
			return nil
		}
		copyFile(sourcePath, dest)
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
