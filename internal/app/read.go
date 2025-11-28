package app

import (
	"fmt"
	"io"
	"log"
	"os"
)

func openFiles(args []string) ([]*os.File, error) {
	files := make([]*os.File, 0, len(args))
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("open file %s: %w", filename, err)
		}

		files = append(files, file)
	}

	return files, nil
}

func closeFiles(files []*os.File) {
	for _, file := range files {
		if err := file.Close(); err != nil {
			log.Fatalln("close file: %w", err)
		}
	}
}

func getReader(files []*os.File) io.Reader {
	readers := make([]io.Reader, 0, len(files))
	for _, file := range files {
		readers = append(readers, file)
	}

	return io.MultiReader(readers...)
}
