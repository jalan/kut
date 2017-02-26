package main

import (
	"github.com/jalan/kut"
	"reflect"
	"testing"
)

var parseToColNumTests = []struct {
	input  string
	output int
}{
	{"1", 1},
	{"123", 123},
}

var parseToColNumInvalidInputs = []string{
	"",
	"+1",
	"-",
	"-7",
	"0",
	"0xAA",
	"1 2",
	"1.0",
	"B",
	"a",
}

func TestParseToColNum(t *testing.T) {
	for _, pair := range parseToColNumTests {
		output, err := parseToColNum(pair.input)
		if err != nil {
			t.Errorf("parseToColNum(%#v): caused error but should not", pair.input)
		}
		if output != pair.output {
			t.Errorf("parseToColNum(%#v): expected %#v but got %#v", pair.input, pair.output, output)
		}
	}

	for _, input := range parseToColNumInvalidInputs {
		_, err := parseToColNum(input)
		if err == nil {
			t.Errorf("parseToColNum(%#v): should cause error but did not", input)
		}
	}
}

var parseToColRangeTests = []struct {
	input  string
	output kut.ColRange
}{
	{"-10", kut.ColRange{Start: 1, End: 10}},
	{"1", kut.ColRange{Start: 1, End: 1}},
	{"100-", kut.ColRange{Start: 100, End: kut.EOL}},
	{"12-12", kut.ColRange{Start: 12, End: 12}},
	{"2-5", kut.ColRange{Start: 2, End: 5}},
}

var parseToColRangeInvalidInputs = []string{
	"",
	"+1",
	"-",
	"-1-1",
	"0",
	"0xAA",
	"1 2",
	"1.0",
	"10,15",
	"10-9",
	"5-2",
	"B",
	"a",
}

func TestParseToColRange(t *testing.T) {
	for _, pair := range parseToColRangeTests {
		output, err := parseToColRange(pair.input)
		if err != nil {
			t.Errorf("parseToColRange(%#v): caused error but should not", pair.input)
		}
		if output != pair.output {
			t.Errorf("parseToColRange(%#v): expected %#v but got %#v", pair.input, pair.output, output)
		}
	}

	for _, input := range parseToColRangeInvalidInputs {
		_, err := parseToColRange(input)
		if err == nil {
			t.Errorf("parseToColRange(%#v): should cause error but did not", input)
		}
	}
}

var parseToListTests = []struct {
	input  string
	output []kut.ColRange
}{
	{"-1", []kut.ColRange{{Start: 1, End: 1}}},
	{"1", []kut.ColRange{{Start: 1, End: 1}}},
	{"1,1", []kut.ColRange{{Start: 1, End: 1}}},
	{"1,2", []kut.ColRange{{Start: 1, End: 1}, {Start: 2, End: 2}}},
	{"1-", []kut.ColRange{{Start: 1, End: kut.EOL}}},
	{"1-1", []kut.ColRange{{Start: 1, End: 1}}},
	{"2,5-9,12-", []kut.ColRange{{Start: 2, End: 2}, {Start: 5, End: 9}, {Start: 12, End: kut.EOL}}},
	{"5-9", []kut.ColRange{{Start: 5, End: 9}}},
}

var parseToListInvalidInputs = []string{
	"",
	"+1",
	"-",
	"-1-1",
	"0",
	"0xAA",
	"1 2",
	"1.0",
	"10-9",
	"5-2",
	"B",
	"a",
}

func TestParseToList(t *testing.T) {
	for _, pair := range parseToListTests {
		output, err := parseToList(pair.input)
		if err != nil {
			t.Errorf("parseToList(%#v): caused error but should not", pair.input)
		}
		if !reflect.DeepEqual(output, pair.output) {
			t.Errorf("parseToList(%#v): expected %#v but got %#v", pair.input, pair.output, output)
		}
	}

	for _, input := range parseToListInvalidInputs {
		_, err := parseToList(input)
		if err == nil {
			t.Errorf("parseToList(%#v): should cause error but did not", input)
		}
	}
}

var parseArgsTests = []struct {
	input  []string
	output []kut.ColRange
}{
	{[]string{"kut", "5-9"}, []kut.ColRange{{Start: 5, End: 9}}},
}

var parseArgsInvalidInputs = [][]string{
	{"kut", "1-3", "7"},
	{"kut", "file"},
	{"kut"},
}

func TestParseArgs(t *testing.T) {
	for _, pair := range parseArgsTests {
		output, err := parseArgs(pair.input)
		if err != nil {
			t.Errorf("parseArgs(%#v): caused error but should not", pair.input)
		}
		if !reflect.DeepEqual(output, pair.output) {
			t.Errorf("parseArgs(%#v): expected %#v but got %#v", pair.input, pair.output, output)
		}
	}

	for _, input := range parseArgsInvalidInputs {
		_, err := parseArgs(input)
		if err == nil {
			t.Errorf("parseArgs(%#v): should cause error but did not", input)
		}
	}
}
