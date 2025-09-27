package grepper

import (
	"bufio"
	"fmt"
	"os"
	"wbgrep/internal/parser"
)

type Grepper struct {
	flags   *parser.Flags
	pattern string
	files   []string
}

func New(flags *parser.Flags, pattern string, files []string) *Grepper {
	return &Grepper{
		flags:   flags,
		pattern: pattern,
		files:   files,
	}
}

func (g *Grepper) Grep() error {
	if len(g.files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		res, err := g.processLines(scanner)

		if err != nil {
			return fmt.Errorf("grep stdin error: %w", err)
		}
		g.outputStream(res)
	} else {
		isOne := true
		if len(g.files) > 1 {
			isOne = false
		}
		for _, f := range g.files {
			res, err := g.processFile(f)
			if err != nil {
				return fmt.Errorf("grep file '%s' error: :%w", f, err)
			}
			g.outputFile(f, res, isOne)
		}
	}
	return nil

}
