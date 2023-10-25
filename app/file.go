package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (app *CommandLineApp) copyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	err = destinationFile.Sync()
	if err != nil {
		return fmt.Errorf("error syncing file: %v", err)
	}

	return nil
}

func (app *CommandLineApp) removeFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (app *CommandLineApp) createFile(fileName string) error {
	filePath := filepath.Join(app.CurrentDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func findFiles(path, targetName string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == targetName || strings.Contains(info.Name(), targetName) {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("error searching for files:", err)
	}
}
