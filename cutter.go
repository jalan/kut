// Package kut provides column cutting for CSV (RFC 4180) files.
package kut

import (
	"bufio"
	"encoding/csv"
	"io"
)

// A ColRange specifies a range of columns to include. Column numbers begin at
// 1. Both Start and End are inclusive.
type ColRange struct {
	Start int
	End   int
}

// EOL can be used as the End in a ColRange to include all remaining columns.
const EOL = int(^uint(0) >> 1)

const lf = byte('\n')
const cr = byte('\r')

// A Cutter reads from an input CSV file and writes only the specified columns
// to an output file.
type Cutter struct {
	buf    *bufio.Reader
	i      *csv.Reader
	o      *csv.Writer
	Ranges []ColRange
}

// NewCutter returns a Cutter that reads from r and writes to w.
func NewCutter(r io.Reader, w io.Writer) *Cutter {
	bufReader := bufio.NewReader(r)
	c := &Cutter{
		buf: bufReader,
		i:   csv.NewReader(bufReader),
		o:   csv.NewWriter(w),
	}
	c.i.FieldsPerRecord = -1
	c.i.LazyQuotes = true
	return c
}

func (c *Cutter) isIncluded(colNum int) bool {
	for _, r := range c.Ranges {
		if colNum >= r.Start && colNum <= r.End {
			return true
		}
	}
	return false
}

// Scan advances one record on the input, outputting only the columns specified
// in Ranges. If there is no input left to read, Scan returns io.EOF.
func (c *Cutter) Scan() error {
	err := c.scan()
	if err != nil {
		return err
	}
	return c.flush()
}

// This is like csv.Reader.Read, except it doesn't skip blank-line records,
// instead returning a nil slice in that case.
func (c *Cutter) read() ([]string, error) {
	bytes, err := c.buf.Peek(1)
	if err != nil {
		return nil, err
	}
	if bytes[0] == lf {
		c.buf.Discard(1)
		return nil, nil
	}
	bytes, err = c.buf.Peek(2)
	if err != nil {
		return nil, err
	}
	if bytes[0] == cr && bytes[1] == lf {
		c.buf.Discard(2)
		return nil, nil
	}
	return c.i.Read()
}

func (c *Cutter) scan() error {
	inputRecord, err := c.read()
	if err != nil {
		return err
	}
	var outputRecord []string
	for i, value := range inputRecord {
		colNum := i + 1 // column numbers begin at 1
		if c.isIncluded(colNum) {
			outputRecord = append(outputRecord, value)
		}
	}
	return c.o.Write(outputRecord)
}

func (c *Cutter) flush() error {
	c.o.Flush()
	return c.o.Error()
}

// ScanAll advances to the end of the input, outputting only the columns
// specified in Ranges. Because ScanAll deliberately reads until EOF, it does
// not report EOF as an error.
func (c *Cutter) ScanAll() error {
	for {
		err := c.scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return c.flush()
}

// SetDelimiter sets the input and output delimiter, which defaults to a comma.
func (c *Cutter) SetDelimiter(d rune) {
	c.i.Comma = d
	c.o.Comma = d
}
