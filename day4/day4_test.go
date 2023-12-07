package main

import (
	"testing"
	"fmt"
)


func TestCheck2(t *testing.T) {
	cases := []struct { line string; expected int } {
		{ "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 4 },
	}

	for _, c := range cases {

		n := Check1(c.line)
		if n != c.expected {
			t.Log(fmt.Sprintf("Error: should be %d, but got %d", c.expected, n))
			t.Fail()
		}
	}
}
