package main

import (
	"fmt"
	"github.com/jalan/kut"
	"os"
	"strconv"
	"strings"
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
			return cr, err
		}
	}

	if len(parts) == 1 {
		cr.End = cr.Start
	} else if parts[1] == "" {
		cr.End = kut.EOL
	} else {
		cr.End, err = parseToColNum(parts[1])
		if err != nil {
			return cr, err
		}
	}

	// Other invalid cases not caught above
	if s == "" || s == "-" || cr.Start > cr.End {
		return cr, fmt.Errorf("invalid column range: %v", s)
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

func parseArgs(args []string) ([]kut.ColRange, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid number of arguments")
	}
	crs, err := parseToList(args[1])
	if err != nil {
		return nil, err
	}
	return crs, nil
}

func die(err error) {
	fmt.Fprintln(os.Stderr, "kut:", err)
	os.Exit(1)
}

func main() {
	crs, err := parseArgs(os.Args)
	if err != nil {
		die(err)
	}
	c := kut.NewCutter(os.Stdin, os.Stdout)
	c.Ranges = crs
	err = c.ScanAll()
	if err != nil {
		die(err)
	}
}
