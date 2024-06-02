package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	log.Println(stdBuffer.String())

}
