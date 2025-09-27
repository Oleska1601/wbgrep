package grepper

import "fmt"

//isMatch bool
func (g *Grepper) outputStream(lines []ResultLine) {

	if g.flags.FlagC {
		fmt.Println(len(lines))
	} else {
		if g.flags.FlagN {
			for _, line := range lines {
				fmt.Printf("%d:%s\n", line.Num, line.Line)
			}
		} else {
			for _, line := range lines {
				fmt.Printf("%s\n", line.Line)
			}
		}
	}

}

// см можно ли упростить!!!!!!!!

func (g *Grepper) outputFile(filename string, lines []ResultLine, isOne bool) {
	if g.flags.FlagC {
		fmt.Println(len(lines))
	} else {
		if !isOne {
			if g.flags.FlagN {
				for _, line := range lines {
					fmt.Printf("%s:%d:%s\n", filename, line.Num, line.Line)
				}
			} else {
				for _, line := range lines {
					fmt.Printf("%s:%s\n", filename, line.Line)
				}
			}
		} else {
			if g.flags.FlagN {
				for _, line := range lines {
					fmt.Printf("%d:%s\n", line.Num, line.Line)
				}
			} else {
				for _, line := range lines {
					fmt.Printf("%s\n", line.Line)
				}
			}
		}

	}
}
