package parser

import (
	"errors"
	"flag"
)

type Flags struct {
	FlagAN int
	FlagBN int
	FlagCN int
	FlagC  bool
	FlagI  bool
	FlagV  bool
	FlagF  bool
	FlagN  bool
}

func Parse() (*Flags, string, []string, error) {
	flags := &Flags{}
	flag.IntVar(&flags.FlagAN, "A", 0, "print additional N lines after each found line")
	flag.IntVar(&flags.FlagBN, "B", 0, "print additional N lines before each found line")
	flag.IntVar(&flags.FlagCN, "C", 0, "print additional N lines before and after each found line")
	flag.BoolVar(&flags.FlagC, "c", false, "show only the number of lines that match the template")
	flag.BoolVar(&flags.FlagI, "i", false, "ignore the register")
	flag.BoolVar(&flags.FlagV, "v", false, "invert the filter")
	flag.BoolVar(&flags.FlagF, "F", false, "treat the template as a fixed string")
	flag.BoolVar(&flags.FlagN, "n", false, "show the line number before each found line")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return nil, "", nil, errors.New("no pattern was provided")
	}

	pattern := args[0]
	files := args[1:]
	return flags, pattern, files, nil
}

//
