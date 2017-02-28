package main

import (
	"github.com/jalan/kut"
	"reflect"
	"testing"
)

var parseToColNumTests = []struct {
	input   string
	output  int
	wantErr bool
}{
	{"1", 1, false},
	{"123", 123, false},
	{"", 0, true},
	{"+1", 0, true},
	{"-", 0, true},
	{"-7", 0, true},
	{"0", 0, true},
	{"0xAA", 0, true},
	{"1 2", 0, true},
	{"1.0", 0, true},
	{"B", 0, true},
	{"a", 0, true},
}

func TestParseToColNum(t *testing.T) {
	for _, test := range parseToColNumTests {
		output, err := parseToColNum(test.input)
		if test.wantErr && err == nil {
			t.Errorf("parseToColNum(%#v): expected an error but got nil", test.input)
		}
		if !test.wantErr && err != nil {
			t.Errorf("parseToColNum(%#v): expected nil error but got %#v", test.input, err)
		}
		if output != test.output {
			t.Errorf("parseToColNum(%#v): expected output %#v but got %#v", test.input, test.output, output)
		}
	}
}

var parseToColRangeTests = []struct {
	input   string
	output  kut.ColRange
	wantErr bool
}{
	{"-10", kut.ColRange{Start: 1, End: 10}, false},
	{"1", kut.ColRange{Start: 1, End: 1}, false},
	{"100-", kut.ColRange{Start: 100, End: kut.EOL}, false},
	{"12-12", kut.ColRange{Start: 12, End: 12}, false},
	{"2-5", kut.ColRange{Start: 2, End: 5}, false},
	{"", kut.ColRange{}, true},
	{"+1", kut.ColRange{}, true},
	{"-", kut.ColRange{}, true},
	{"-1-1", kut.ColRange{}, true},
	{"0", kut.ColRange{}, true},
	{"0xAA", kut.ColRange{}, true},
	{"1 2", kut.ColRange{}, true},
	{"1.0", kut.ColRange{}, true},
	{"10,15", kut.ColRange{}, true},
	{"10-9", kut.ColRange{}, true},
	{"5-2", kut.ColRange{}, true},
	{"B", kut.ColRange{}, true},
	{"a", kut.ColRange{}, true},
}

func TestParseToColRange(t *testing.T) {
	for _, test := range parseToColRangeTests {
		output, err := parseToColRange(test.input)
		if test.wantErr && err == nil {
			t.Errorf("parseToColRange(%#v): expected an error but got nil", test.input)
		}
		if !test.wantErr && err != nil {
			t.Errorf("parseToColRange(%#v): expected nil error but got %#v", test.input, err)
		}
		if output != test.output {
			t.Errorf("parseToColRange(%#v): expected output %#v but got %#v", test.input, test.output, output)
		}
	}
}

var parseToListTests = []struct {
	input   string
	output  []kut.ColRange
	wantErr bool
}{
	{"-1", []kut.ColRange{{Start: 1, End: 1}}, false},
	{"1", []kut.ColRange{{Start: 1, End: 1}}, false},
	{"1,1", []kut.ColRange{{Start: 1, End: 1}}, false},
	{"1,2", []kut.ColRange{{Start: 1, End: 1}, {Start: 2, End: 2}}, false},
	{"1-", []kut.ColRange{{Start: 1, End: kut.EOL}}, false},
	{"1-1", []kut.ColRange{{Start: 1, End: 1}}, false},
	{"2,5-9,12-", []kut.ColRange{{Start: 2, End: 2}, {Start: 5, End: 9}, {Start: 12, End: kut.EOL}}, false},
	{"5-9", []kut.ColRange{{Start: 5, End: 9}}, false},
	{"", nil, true},
	{"+1", nil, true},
	{"-", nil, true},
	{"-1-1", nil, true},
	{"0", nil, true},
	{"0xAA", nil, true},
	{"1 2", nil, true},
	{"1.0", nil, true},
	{"10-9", nil, true},
	{"5-2", nil, true},
	{"B", nil, true},
	{"a", nil, true},
}

func TestParseToList(t *testing.T) {
	for _, test := range parseToListTests {
		output, err := parseToList(test.input)
		if test.wantErr && err == nil {
			t.Errorf("parseToList(%#v): expected an error but got nil", test.input)
		}
		if !test.wantErr && err != nil {
			t.Errorf("parseToList(%#v): expected nil error but got %#v", test.input, err)
		}
		if !reflect.DeepEqual(output, test.output) {
			t.Errorf("parseToList(%#v): expected output %#v but got %#v", test.input, test.output, output)
		}
	}
}

var parseArgsTests = []struct {
	input   []string
	output  []kut.ColRange
	wantErr bool
}{
	{[]string{"kut", "5-9"}, []kut.ColRange{{Start: 5, End: 9}}, false},
	{[]string{"kut", "1-3", "7"}, nil, true},
	{[]string{"kut", "file"}, nil, true},
	{[]string{"kut"}, nil, true},
}

func TestParseArgs(t *testing.T) {
	for _, test := range parseArgsTests {
		output, err := parseArgs(test.input)
		if test.wantErr && err == nil {
			t.Errorf("parseArgs(%#v): expected an error but got nil", test.input)
		}
		if !test.wantErr && err != nil {
			t.Errorf("parseArgs(%#v): expected nil error but got %#v", test.input, err)
		}
		if !reflect.DeepEqual(output, test.output) {
			t.Errorf("parseArgs(%#v): expected output %#v but got %#v", test.input, test.output, output)
		}
	}
}
