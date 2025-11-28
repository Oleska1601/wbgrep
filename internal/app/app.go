package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/Oleska1601/wbgrep/internal/grepper"
	"github.com/Oleska1601/wbgrep/internal/parser"
)

// Run executes the cut utility with provided flags and files.
// It processes input from files or stdin and writes results to stdout.
func Run(flags *parser.Flags, pattern string, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		cancel() // cancel when Ctrl+C
	}()

	var reader io.Reader
	if len(args) == 0 {
		reader = os.Stdin
	} else {
		files, err := openFiles(args)
		defer closeFiles(files)
		if err != nil {
			log.Fatalln(fmt.Errorf("openFiles: %w", err))
		}

		reader = getReader(files)
	}

	grepper := grepper.New(flags, pattern)
	outputLines := grepper.Grep(ctx, reader)
	for line := range outputLines {
		fmt.Println(line)
	}

}
