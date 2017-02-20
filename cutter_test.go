package kut

import (
	"bytes"
	"io"
	"testing"
)

func TestNewCutter(t *testing.T) {
	i := new(bytes.Buffer)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	if c == nil {
		t.Errorf("NewCutter returned nil")
	}
}

var testScanInput string = `aaa,bbb,ccc
123,456,789`

func TestScan(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	err := c.Scan()
	if err != nil {
		t.Errorf("Scan returned an error on a valid record")
	}
}

func TestScanEOF(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	var err error
	for err == nil {
		err = c.Scan()
	}
	if err != io.EOF {
		t.Errorf("Scan should eventually return io.EOF")
	}
}

func TestScanNoRanges(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	c.Scan()
	if o.String() != "\n" {
		t.Errorf("Scan with nil Ranges should output one empty record")
	}
}

func TestScanFullRange(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	c.Ranges = append(c.Ranges, ColRange{1, EOL})
	c.Scan()
	if o.String() != "aaa,bbb,ccc\n" {
		t.Errorf("Scan with Ranges [{1, EOF}] should output one full record")
	}
}

func TestScanOneColumn(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	c.Ranges = append(c.Ranges, ColRange{2, 2})
	c.Scan()
	if o.String() != "bbb\n" {
		t.Errorf("Scan with Ranges %v should output one column", c.Ranges)
	}
}

func TestScanTwoColumns(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	c.Ranges = append(c.Ranges, ColRange{2, 3})
	c.Scan()
	if o.String() != "bbb,ccc\n" {
		t.Errorf("Scan with Ranges %v should output two columns", c.Ranges)
	}
}

func TestScanMultipleRanges(t *testing.T) {
	i := bytes.NewBufferString(testScanInput)
	o := new(bytes.Buffer)
	c := NewCutter(i, o)
	c.Ranges = append(c.Ranges, ColRange{1, 1})
	c.Ranges = append(c.Ranges, ColRange{3, 3})
	c.Scan()
	if o.String() != "aaa,ccc\n" {
		t.Errorf("Scan with Ranges %v should output two columns", c.Ranges)
	}
}
