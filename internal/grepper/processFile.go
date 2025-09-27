package grepper

import (
	"bufio"
	"fmt"
	"os"
)

func (g *Grepper) processFile(filename string) ([]ResultLine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file '%s' error: :%w", filename, err)
	}
	scanner := bufio.NewScanner(file)
	res, err := g.processLines(scanner)
	if err != nil {
		return nil, err
	}
	return res, nil
}
