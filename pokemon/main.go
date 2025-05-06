package main

import (
	"fmt"
	"os/exec"
	"pokemon/app"
)

func main() {
	app.Run()
}

func openHTMLFile() {
	filePath := "ui/index.html"

	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/C", "start", filePath)

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
}
