package main

import (
	"log"

	"github.com/Oleska1601/wbgrep/internal/app"
	"github.com/Oleska1601/wbgrep/internal/parser"
)

func main() {
	flags, pattern, files, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(flags, pattern, files)
}
