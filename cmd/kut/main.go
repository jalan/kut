package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jalan/kut"
)

const help = `Usage: kut [OPTION] LIST

kut is a cut command for CSV (RFC 4180) input. kut reads from standard input
and writes to standard output.

LIST is a list of ranges separated by commas, using the same rules as cut(1).

Options:
  -d, --delimiter=DELIM    Use DELIM as field delimiters instead of commas.
  -h, --help               Show this help message.
`

func main() {
	delim, crs, err := parseArgs(os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Print(help)
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "kut:", err)
		os.Exit(1)
	}
	c := kut.NewCutter(os.Stdin, os.Stdout)
	c.SetDelimiter(delim)
	c.Ranges = crs
	err = c.ScanAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "kut:", err)
		os.Exit(2)
	}
}
