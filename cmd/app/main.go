package main

import (
	"log"
	"wbgrep/internal/grepper"
	"wbgrep/internal/parser"
)

func main() {
	flags, pattern, files, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	grepper := grepper.New(flags, pattern, files)
	err = grepper.Grep()
	if err != nil {
		log.Fatalln(err)
	}

}
