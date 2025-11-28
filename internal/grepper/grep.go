package grepper

import (
	"bufio"
	"context"
	"io"

	"github.com/Oleska1601/wbgrep/internal/parser"
)

// Grepper provides functionality for searching text patterns in input streams
type Grepper struct {
	flags   *parser.Flags
	pattern string
}

// New creates and returns a new Grepper instance with the specified flags and pattern
func New(flags *parser.Flags, pattern string) *Grepper {
	return &Grepper{
		flags:   flags,
		pattern: pattern,
	}
}

// Grep performs pattern matching on the provided reader and returns a channel
// where matching lines would be written. The search could be cancelled due to ctrl+c and context
func (g *Grepper) Grep(ctx context.Context, reader io.Reader) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		scanner := bufio.NewScanner(reader)
		if g.flags.FlagC {
			g.processFlagC(ctx, scanner, output)
			return
		}

		g.processLines(ctx, scanner, output)
	}()

	return output
}
