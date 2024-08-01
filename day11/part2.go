package main

import (
	"os"
	"log"
	"slices"
	"fmt"
)


func MoveGalaxies(universe [][]rune, galaxies []Point) []Point {
	var newGalaxies []Point = make([]Point, len(galaxies))
	copy(newGalaxies, galaxies)

	for y, line := range(universe) {
		if !slices.Contains(line, '#') {
			for idx := range(galaxies) {
				if galaxies[idx].y > y {
					newGalaxies[idx].y += 999999
				}
			}
		}
	}

	for x := 0; x < len(universe[0]); x++ {
		empty := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x] == '#' {
				empty = false
			}
		}

		if empty {
			for idx := range(galaxies) {
				if galaxies[idx].x > x {
					newGalaxies[idx].x += 999999
				}
			}
		}
	}


	return newGalaxies
}


func Part2() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	universe := Parse(string(f))
	galaxies := FindGalaxies(universe)
	fmt.Println(galaxies)
	galaxies = MoveGalaxies(universe, galaxies)
	fmt.Println(galaxies, SumOfDistances(universe, galaxies))
}
