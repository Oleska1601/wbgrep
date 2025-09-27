package grepper

import (
	"bufio"
	"fmt"
	"math"
)

type ResultLine struct {
	Line string
	Num  int
}

func newResultLine(line string, num int) ResultLine {
	return ResultLine{
		Line: line,
		Num:  num,
	}
}

func (g *Grepper) processLines(scanner *bufio.Scanner) ([]ResultLine, error) {

	size := int(math.Max(float64(g.flags.FlagBN), float64(g.flags.FlagCN)))
	queueBefore := newQueue(size + 1)

	end := int(math.Max(float64(g.flags.FlagAN), float64(g.flags.FlagCN)))
	var afterCount int
	lineNum := 0
	var res []ResultLine

	for scanner.Scan() {
		line := scanner.Text()

		lineNum++
		lineResult := newResultLine(line, lineNum)
		if afterCount > 0 {
			res = append(res, lineResult)
			afterCount--
		} else if afterCount == 0 {
			queueBefore.enqueue(lineResult)
		}
		if g.isMatch(line) {
			values := queueBefore.getAll()
			res = append(res, values...)
			queueBefore.clear()
			afterCount = end
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan string error: %w", err)
	}

	return res, nil
}
