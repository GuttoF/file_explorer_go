package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (app *CommandLineApp) changeDirectory(path string) error {
	if path == ".." {
		app.CurrentDir = filepath.Dir(app.CurrentDir)
		return nil
	}

	newDir := filepath.Join(app.CurrentDir, path)

	fileInfo, err := os.Stat(newDir)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("path is not a directory")
	}

	app.CurrentDir = newDir
	return nil
}

func listFilesInDirectory(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (app *CommandLineApp) createDirectory(directoryPath string) error {
	err := os.Mkdir(directoryPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
