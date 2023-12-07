package main

import (
	"testing"
	"fmt"
	"github.com/google/go-cmp/cmp"
)


func TestCheck2(t *testing.T) {
	cases := []struct { line, prev, next string; expected map[Vertex][]int } {
		{ "467..114..", "..........", "...*......", map[Vertex][]int{ {1,3}: []int{467} } },
		{ "...*......", "467..114..", "..35..633.", map[Vertex][]int{  } },
		{ "..35..633.", "...*......", "......#...", map[Vertex][]int{ {-1,3}: []int{35} } },
		{ "......#...", "..35..633.", "617*......", map[Vertex][]int{  } },
		{ "617*......", "......#...", ".....+.58.", map[Vertex][]int{ {0,3}: []int{617} } },
		{ ".....+.58.", "617*......", "..592.....", map[Vertex][]int{  } },
		{ "..592.....", ".....+.58.", "......755.", map[Vertex][]int{  } },
		{ "......755.", "..592.....", "...$.*....", map[Vertex][]int{ {1,5}: []int{755} } },
		{ "...$.*....", "......755.", ".664.598..", map[Vertex][]int{  } },
		{ ".664.598..", "...$.*....", "..........", map[Vertex][]int{ {-1,5}: []int{598} } },
	}

	for _, c := range cases {

		m := Check2(c.line, c.prev, c.next)
		if !cmp.Equal(m, c.expected) {
			t.Log(fmt.Sprintf("Error: should be %d, but got %q", c.expected, m))
			t.Fail()
		}
	}
}
