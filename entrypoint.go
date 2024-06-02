package main

import (
	"bytes"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	arguments = append(arguments, args[7])

	cmd := exec.Command("migrate", arguments...)
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
