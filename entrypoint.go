package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args

	arguments := args[1]

	cmd := exec.Command("migrate", arguments)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Printf("error: %v", err)
		exitError, ok := err.(*exec.ExitError)
		if ok {
			os.Exit(exitError.ExitCode())
		}
	}

	log.Println(stdBuffer.String())

}
