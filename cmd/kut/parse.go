package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jalan/kut"
)

func parseToColNum(s string) (int, error) {
	colNum, err := strconv.Atoi(s)
	if err != nil || s[0] == '+' || s[0] == '-' || colNum == 0 {
		return 0, fmt.Errorf("invalid column number: %v", s)
	}
	return colNum, nil
}

func parseToColRange(s string) (kut.ColRange, error) {
	var cr kut.ColRange
	var err error

	parts := strings.SplitN(s, "-", 2)

	if parts[0] == "" {
		cr.Start = 1
	} else {
		cr.Start, err = parseToColNum(parts[0])
		if err != nil {
			return kut.ColRange{}, err
		}
	}

	if len(parts) == 1 {
		cr.End = cr.Start
	} else if parts[1] == "" {
		cr.End = kut.EOL
	} else {
		cr.End, err = parseToColNum(parts[1])
		if err != nil {
			return kut.ColRange{}, err
		}
	}

	// Other invalid cases not caught above
	if s == "" || s == "-" || cr.Start > cr.End {
		return kut.ColRange{}, fmt.Errorf("invalid column range: %v", s)
	}

	return cr, nil
}

func contains(crs []kut.ColRange, cr kut.ColRange) bool {
	for _, elem := range crs {
		if elem == cr {
			return true
		}
	}
	return false
}

func parseToList(s string) ([]kut.ColRange, error) {
	var crs []kut.ColRange
	for _, input := range strings.Split(s, ",") {
		cr, err := parseToColRange(input)
		if err != nil {
			return nil, err
		}
		if !contains(crs, cr) {
			crs = append(crs, cr)
		}
	}
	return crs, nil
}

func stringToDelimiter(s string) (rune, error) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, fmt.Errorf("delimiter must be a single rune")
	}
	d, _ := utf8.DecodeRuneInString(s)
	if d == utf8.RuneError {
		return 0, fmt.Errorf("delimiter must be valid UTF-8")
	}
	return d, nil
}

func parseArgs(args []string) (rune, []kut.ColRange, error) {
	var delimStr string
	flagSet := flag.NewFlagSet("kut", flag.ContinueOnError)
	flagSet.SetOutput(ioutil.Discard)
	flagSet.StringVar(&delimStr, "d", ",", "")
	flagSet.StringVar(&delimStr, "delimiter", ",", "")
	err := flagSet.Parse(args)
	if err == flag.ErrHelp {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, fmt.Errorf("invalid option")
	}
	delim, err := stringToDelimiter(delimStr)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid delimiter")
	}
	if len(flagSet.Args()) != 1 {
		return 0, nil, fmt.Errorf("invalid number of arguments")
	}
	crs, err := parseToList(flagSet.Arg(0))
	if err != nil {
		return 0, nil, err
	}
	return delim, crs, nil
}
