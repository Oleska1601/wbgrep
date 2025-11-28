package grepper

import (
	"bufio"
	"context"
	"fmt"
	"log"
)

func (g *Grepper) formatLine(line string, num int) string {
	if g.flags.FlagN {
		return fmt.Sprintf("%d:%s", num, line)
	}
	return line
}

func (g *Grepper) processLines(ctx context.Context, scanner *bufio.Scanner, output chan<- string) {
	bufSize := max(g.flags.FlagBN, g.flags.FlagCN)
	buffer := newBuffer(bufSize + 1) // общее кол-во значений, которые выводятся до найденной строки +1 для хранения самой найденной строки и последующего корректного вывода
	linesAfter := max(g.flags.FlagAN, g.flags.FlagCN)

	var afterCount int
	var lineNum int

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if afterCount > 0 {
			select {
			case <-ctx.Done():
				return
			case output <- g.formatLine(line, lineNum):
				afterCount--
			}
		} else {
			buffer.add(g.formatLine(line, lineNum))
		}

		if g.isMatch(line) {
			for _, bufLine := range buffer.getAll() {
				// check
				select {
				case <-ctx.Done():
					return
				case output <- bufLine:
				}
			}

			buffer.clear()
			afterCount = linesAfter
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("scanner error: %v\n", err)
	}
}
