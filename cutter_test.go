package kut

import (
	"bytes"
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
