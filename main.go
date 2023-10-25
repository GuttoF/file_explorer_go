package main

import (
	"file_explorer_go/app"
)

func main() {
	myApp := app.NewCommandLineApp()
	myApp.Run()
}
