package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CommandLineApp represents the state of the command-line application.
type CommandLineApp struct {
	CurrentDir string
}

// NewCommandLineApp creates a new CommandLineApp instance and initializes it with the current working directory.
func NewCommandLineApp() *CommandLineApp {
	return &CommandLineApp{
		CurrentDir: getCurrentDir(),
	}
}

// Run is the main application loop that handles user input and executes commands.
func (app *CommandLineApp) Run() {
	for {
		fmt.Printf("[%s] $ ", app.CurrentDir)

		var command string
		_, err := fmt.Scan(&command)
		if err != nil {
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
			err = app.changeDirectory(path)
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
			err = app.copyFile(source, destination)
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
			err = app.removeFile(path)
			if err != nil {
				fmt.Println("error removing file", err)
			}

		case "ls":
			listFilesInDirectory(app.CurrentDir)

		case "mkdir":
			var dirName string
			_, err := fmt.Scan(&dirName)
			if err != nil {
				fmt.Println("error reading directory name", err)
				continue
			}
			directoryPath := filepath.Join(app.CurrentDir, dirName)
			err = app.createDirectory(directoryPath)
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
			findFiles(app.CurrentDir, targetFileName)

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
			err = app.createFile(fileName)
			if err != nil {
				fmt.Println("error creating file", err)
			}

		case "vim":
			err := openVimEditor()
			if err != nil {
				fmt.Println("error opening Vim:", err)
			}

		case "pwd":
			app.printWorkingDirectory()

		default:
			fmt.Println("unknown command", command)
		}
	}
}

// getCurrentDir gets the current working directory.
func getCurrentDir() string {
	currentDir, _ := os.Getwd()
	return currentDir
}
