package grepper

import (
	"regexp"
	"strings"
)

func (g *Grepper) isMatch(line string) bool {
	isMatch := false
	processLine := line
	processPattern := g.pattern
	if g.flags.FlagI {
		processLine = strings.ToLower(line)
		processPattern = strings.ToLower(g.pattern)
	}

	if g.flags.FlagF {
		isMatch = processLine == processPattern
	} else {
		re := regexp.MustCompile(processPattern)
		isMatch = re.MatchString(processLine)
	}

	if g.flags.FlagV {
		return !isMatch
	}

	return isMatch
}
