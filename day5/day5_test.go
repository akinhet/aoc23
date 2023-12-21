package main

import (
	"testing"
	"fmt"
)


func TestConvertUsingMap(t *testing.T) {
	cases := []struct { num int; m []Map; expected int } {
		{ 79, []Map{{50, 98, 2}, {52, 50, 48}}, 81},
	}

	for _, c := range cases {

		n := ConvertUsingMap(c.num, c.m)
		if n != c.expected {
			t.Log(fmt.Sprintf("Error: should be %d, but got %d", c.expected, n))
			t.Fail()
		}
	}
}
