package kut

import (
	"bytes"
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
	{"\n", []ColRange{}, io.EOF, ""},
	{"abc", []ColRange{{1, 1}}, nil, "abc\n"},
	{"abc,def,ghi", []ColRange{{1, 1}, {3, 3}}, nil, "abc,ghi\n"},
	{"abc,def,ghi", []ColRange{{1, EOL}}, nil, "abc,def,ghi\n"},
	{"abc,def,ghi", []ColRange{{2, 2}}, nil, "def\n"},
	{"abc,def,ghi", []ColRange{{2, 3}}, nil, "def,ghi\n"},
	{"abc,def,ghi", []ColRange{{3, 2}}, nil, "\n"},
	{"abc,def,ghi", []ColRange{{3, EOL}}, nil, "ghi\n"},
	{"abc,def,ghi", []ColRange{{4, EOL}}, nil, "\n"},
	{"abc,def,ghi", []ColRange{}, nil, "\n"},
}

func TestScan(t *testing.T) {
	for _, test := range scanTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		if err := c.Scan(); err != test.err {
			t.Errorf("expected return value %#v but got %#v", test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("expected output %#v but got %#v", test.output, outString)

		}
	}
}

var scanAllTests = []struct {
	input  string
	ranges []ColRange
	err    error
	output string
}{
	{"abc,def,ghi\njkl,mno,pqr\n", []ColRange{{2, 2}}, nil, "def\nmno\n"},
}

func TestScanAll(t *testing.T) {
	for _, test := range scanAllTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		if err := c.ScanAll(); err != test.err {
			t.Errorf("expected return value %#v but got %#v", test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("expected output %#v but got %#v", test.output, outString)

		}
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
	for _, test := range setDelimiterTests {
		outBuf := new(bytes.Buffer)
		c := NewCutter(bytes.NewBufferString(test.input), outBuf)
		c.Ranges = test.ranges
		c.SetDelimiter(test.delimiter)
		if err := c.ScanAll(); err != test.err {
			t.Errorf("expected return value %#v but got %#v", test.err, err)
		}
		if outString := outBuf.String(); outString != test.output {
			t.Errorf("expected output %#v but got %#v", test.output, outString)

		}
	}
}
