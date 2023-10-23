package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// First directory
	currentDir, _ := os.Getwd()

	for {
		fmt.Printf("[%s] $ ", currentDir)

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

		case "clear":
			clearTerminal()

		case "find":
			var targetFileName string
			_, err := fmt.Scan(&targetFileName)
			if err != nil {
				fmt.Println("error reading target file name", err)
				continue
			}
			findFiles(currentDir, targetFileName)

		case "exit":
			fmt.Println("exiting...")
			os.Exit(0)

		case "touch":
			var fileName string
			_, err := fmt.Scan(&fileName)
			if err != nil {
				fmt.Println("error reading file name", err)
				continue
			}
			err = createFile(currentDir, fileName)
			if err != nil {
				fmt.Println("error creating file", err)
			}

		case "vim":
			err := openVimEditor()
			if err != nil {
				fmt.Println("error opening Vim:", err)
			}

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

func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear") // "clear" in Linux.
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // "cls" in Windows.
	default:
		fmt.Println("Clearing the terminal is not supported on this operating system.")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createFile(directory, fileName string) error {
	filePath := filepath.Join(directory, fileName)
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

func openVimEditor() error {
	cmd := exec.Command("vim")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
