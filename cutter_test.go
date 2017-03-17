package kut

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestNewCutter(t *testing.T) {
	if c := NewCutter(new(bytes.Buffer), new(bytes.Buffer)); c == nil {
		t.Errorf("NewCutter returned nil")
	}
}

var scanTests = []struct {
	input  string
	ranges []ColRange
	err    error
	output string
}{
	{"", []ColRange{}, io.EOF, ""},
	{"\n", []ColRange{}, nil, "\n"},
	{"\r\n", []ColRange{}, nil, "\n"},
	{"abc", []ColRange{{1, 1}}, nil, "abc\n"},
	{"abc,def,ghi", []ColRange{{1, 1}, {3, 3}}, nil, "abc,ghi\n"},
	{"abc,def,ghi", []ColRange{{1, EOL}}, nil, "abc,def,ghi\n"},
	{"abc,def,ghi", []ColRange{{2, 2}}, nil, "def\n"},
	{"abc,def,ghi", []ColRange{{2, 3}}, nil, "def,ghi\n"},
	{"abc,def,ghi", []ColRange{{3, 2}}, nil, "\n"},
	{"abc,def,ghi", []ColRange{{3, EOL}}, nil, "ghi\n"},
	{"abc,def,ghi", []ColRange{{4, EOL}}, nil, "\n"},
	{"abc,def,ghi", []ColRange{}, nil, "\n"},
	{"abc,def,ghi,j\"k", []ColRange{{2, 4}}, nil, "def,ghi,\"j\"\"k\"\n"},
}

func TestScan(t *testing.T) {
	for i, test := range scanTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		if err := c.Scan(); err != test.err {
			t.Errorf("TestScan %v: expected error %#v but got %#v", i, test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("TestScan %v: expected output %#v but got %#v", i, test.output, outString)

		}
	}
}

var scanAllTests = []struct {
	input  string
	ranges []ColRange
	err    error
	output string
}{
	{"1,2,3\n\n\n1,2,3", []ColRange{{2, 2}}, nil, "2\n\n\n2\n"},
	{"\n\n\n", []ColRange{{2, 2}}, nil, "\n\n\n"},
	{"a\na", []ColRange{{1, 5}}, nil, "a\na\n"},
	{"abc,def,ghi\njkl,mn\ro,pqr\n", []ColRange{{2, 2}}, nil, "def\n\"mn\ro\"\n"},
	{"abc,def,ghi\njkl,mno,pqr", []ColRange{{2, 2}}, nil, "def\nmno\n"},
	{"abc,def,ghi\njkl,mno,pqr\n", []ColRange{{2, 2}}, nil, "def\nmno\n"},
	{"abc,def,ghi\r\njkl,mno,pqr", []ColRange{{2, 2}}, nil, "def\nmno\n"},
	{"abc,def,ghi\r\njkl,mno,pqr\r\n", []ColRange{{2, 2}}, nil, "def\nmno\n"},
}

func TestScanAll(t *testing.T) {
	for i, test := range scanAllTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		if err := c.ScanAll(); err != test.err {
			t.Errorf("TestScanAll %v: expected error %#v but got %#v", i, test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("TestScanAll %v: expected output %#v but got %#v", i, test.output, outString)

		}
	}
}

type BrokenReader struct{}

func (br *BrokenReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("this reader never works")
}

func TestScanAllUnknownError(t *testing.T) {
	in := new(BrokenReader)
	out := new(bytes.Buffer)
	c := NewCutter(in, out)
	if err := c.ScanAll(); err == nil {
		t.Errorf("TestScanAllUnknownError: expected an error but got nil")
	}
}

var setDelimiterTests = []struct {
	input     string
	ranges    []ColRange
	delimiter rune
	err       error
	output    string
}{
	{"abc;def,ghi;123\njkl;mno,pqr;45678\n", []ColRange{{2, 2}}, ';', nil, "def,ghi\nmno,pqr\n"},
}

func TestSetDelimiter(t *testing.T) {
	for i, test := range setDelimiterTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		c.SetDelimiter(test.delimiter)
		if err := c.ScanAll(); err != test.err {
			t.Errorf("TestSetDelimiter %v: expected error %#v but got %#v", i, test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("TestSetDelimiter %v: expected output %#v but got %#v", i, test.output, outString)

		}
	}
}

func TestDelimiter(t *testing.T) {
	want := '日'
	buf := new(bytes.Buffer)
	c := NewCutter(buf, buf)
	c.SetDelimiter(want)
	if got := c.Delimiter(); got != want {
		t.Errorf("TestDelimiter: expected %q but got %q", want, got)
	}
}
