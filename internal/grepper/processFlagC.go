package grepper

import (
	"bufio"
	"context"
	"fmt"
	"log"
)

func (g *Grepper) processFlagC(ctx context.Context, scanner *bufio.Scanner, output chan<- string) {
	var counter int
	for scanner.Scan() {
		line := scanner.Text()
		if g.isMatch(line) {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("scanner error: %v\n", err)
		return
	}

	select {
	case <-ctx.Done():
	case output <- fmt.Sprintf("%d", counter):
	}
}
