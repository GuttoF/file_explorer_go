package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// First directory
	currentDir := "."

	for {
		fmt.Printf("[%s] $", currentDir)

		var command string
		_, err := fmt.Scan(&command)
		if err != nil { // Nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
			fmt.Println("error reading the input", err)
			return
		}

		switch command {
		case "cd":
			var path string

			_, err := fmt.Scan(&path)
			if err != nil {
				fmt.Println("error reading path", err)
				continue
			}
			err = changeDirectory(&currentDir, path)
			if err != nil {
				fmt.Println("error changing directory", err)
			}

		case "cp":
			var source, destination string

			_, err := fmt.Scanln(&source, &destination)
			if err != nil {
				fmt.Println("error reading source and destination", err)
				continue
			}
			err = copyFile(source, destination)
			if err != nil {
				fmt.Println("error copying file", err)
			}

		case "rm":
			var path string
			_, err := fmt.Scan(&path)
			if err != nil {
				fmt.Println("error reading path", err)
				continue
			}
			err = removeFile(path)
			if err != nil {
				fmt.Println("error removing file", err)
			}

		case "ls":
			listFilesInDirectory(currentDir)

		case "mkdir":
			var dirName string
			_, err := fmt.Scan(&dirName)
			if err != nil {
				fmt.Println("error reading directory name", err)
				continue
			}
			directoryPath := filepath.Join(currentDir, dirName)
			err = createDirectory(directoryPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("created directory %s\n", directoryPath)

		default:
			fmt.Println("unknown command", command)
		}
	}
}

func changeDirectory(currentDir *string, path string) error {
	if path == ".." {
		*currentDir = filepath.Dir(*currentDir)
		return nil
	}

	newDir := filepath.Join(*currentDir, path)

	fileInfo, err := os.Stat(newDir)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("path is not a directory")
	}

	*currentDir = newDir
	return nil
}

func copyFile(source, destination string) error {
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

func removeFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

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

func createDirectory(directoryPath string) error {
	err := os.Mkdir(directoryPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
