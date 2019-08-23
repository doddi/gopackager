package compress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Folder(folderToZip string, name string) {
	newZipFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	writer := zip.NewWriter(newZipFile)
	defer writer.Close()

	_ = filepath.Walk(folderToZip, func(path string, info os.FileInfo, err error) error {
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
