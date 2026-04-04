package main

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	_, e := parseTime("123000", "20261228")
	if e != nil {
		t.Fatal(e)
	}
}
