package app

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Clearing the terminal is not supported on this operating system.")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func openVimEditor() error {
	cmd := exec.Command("vim")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
