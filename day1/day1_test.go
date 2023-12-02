package main

import (
	"testing"
	"fmt"
)

func TestExtract(t *testing.T) {
	cases := []struct { line string; expected int } {
		{ "two1nine",			29 },
		{ "eightwothree",		83 },
		{ "abcone2threexyz",	13 },
		{ "xtwone3four",		24 },
		{ "4nineeightseven2",	42 },
		{ "zoneight234",		14 },
		{ "7pqrstsixteen",		76 },
		{ "sq5fivetwothree1",	51 },
		{ "six5gc",				65 },
		{ "oneight",			18 },
		{ "treb7uchet",			77 },
		{ "five",				55 },
		{ "threefivethree",		33 },
		{ "eightwo",			82 },
		{ "fiveeight3sppjtccnineeighteightnffgtlsdj",	58 },
		{ "threethreetwothree",	33 },
	}

	for _, c := range cases {

		num := Extract(c.line)
		if num != c.expected {
			t.Log(fmt.Sprintf("Error: should be %d, but got %d", c.expected, num))
			t.Fail()
		}
	}
}
