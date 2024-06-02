package main

import (
	"bytes"
	"log"
)

func main() {
	var stdBuffer bytes.Buffer
	log.Println(stdBuffer.String())

}
