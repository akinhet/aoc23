package main

import (
	"os"
	"log"
	"strings"
	"fmt"
	"slices"
)

type Point struct {
	x, y int
}


func Parse(f string) ([][]rune) {
	var universe [][]rune

	lines := strings.Split(f, "\n")
	lines = lines[:len(lines)-1]
	for y, line := range(lines) {
		universe = append(universe, make([]rune, len(line)))
		for x, char := range(line) {
			universe[y][x] = char
		}
	}

	return universe
}


func ExpandUniverse(universe [][]rune) [][]rune {
	var expandedUniverse [][]rune
	var count int
	for _, line := range(universe) {
		expandedUniverse = append(expandedUniverse, line)
		if !slices.Contains(line, '#') {
			expandedUniverse = append(expandedUniverse, line)
			count++
		}
	}

	for x := 0; x < len(expandedUniverse[0]); x++ {
		empty := true
		for y := 0; y < len(expandedUniverse); y++ {
			if expandedUniverse[y][x] == '#' {
				empty = false
			}
		}

		if empty {
			count++
			for y := 0; y < len(expandedUniverse); y++ {
				expandedUniverse[y] = slices.Insert(expandedUniverse[y], x, '.')
			}
			x++
		}
	}
	fmt.Println("lines added:", count)

	return expandedUniverse
}


func FindGalaxies(universe [][]rune) []Point {
	var galaxies []Point

	for y, line := range(universe) {
		for x, char := range(line) {
			if char == '#' {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}

	return galaxies
}


func SumOfDistances(universe [][]rune, galaxies []Point) int {
	var sum int

	for len(galaxies) > 1 {
		for i, galaxy := range(galaxies) {
			if i == 0 {
				continue
			}
			if max(galaxy.x, galaxies[0].x) == galaxy.x {
				sum += galaxy.x - galaxies[0].x + galaxy.y - galaxies[0].y
			} else {
				sum += galaxies[0].x - galaxy.x + galaxy.y - galaxies[0].y
			}
		}
		galaxies = galaxies[1:]
	}

	return sum
}


func Part1() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	universe := Parse(string(f))
	for _, row := range(universe) {
		for _, c := range(row) {
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	universe = ExpandUniverse(universe)
	galaxies := FindGalaxies(universe)

	for _, row := range(universe) {
		for _, c := range(row) {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
	fmt.Println(galaxies)

	fmt.Println(SumOfDistances(universe, galaxies))
}
