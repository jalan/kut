// Package kut provides column cutting for CSV (RFC 4180) files.
package kut

import (
	"encoding/csv"
	"io"
)

// A ColRange specifies a range of columns to include. Column numbers begin at
// 1. Both Start and End are inclusive.
type ColRange struct {
	Start int
	End   int
}

// A Cutter reads from an input CSV file and writes only the specified columns
// to an output file.
type Cutter struct {
	i      *csv.Reader
	o      *csv.Writer
	Ranges []ColRange
}

// NewCutter returns a Cutter that reads from r and writes to w.
func NewCutter(r io.Reader, w io.Writer) *Cutter {
	c := &Cutter{
		i: csv.NewReader(r),
		o: csv.NewWriter(w),
	}
	c.i.FieldsPerRecord = -1
	return c
}

// Scan advances one record on the input, outputting only the columns specified
// in Ranges. If there is no input left to read, Scan returns io.EOF.
func (c *Cutter) Scan() error {
	return nil
}

// ScanAll advances to the end of the input, outputting only the columns
// specified in Ranges. Because ScanAll deliberately reads until EOF, it does
// not report EOF as an error.
func (c *Cutter) ScanAll() error {
	return nil
}
