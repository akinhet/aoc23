package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)


func Parse(f string) [][]int {
	var sensors [][]int

	lines := strings.Split(f, "\n")
	for i, line := range(lines) {
		if line == "" {
			continue
		}
		sensor := strings.Split(line, " ")
		sensors = append(sensors, make([]int, len(sensor)))

		for j, num := range(sensor) {
			var err error
			sensors[i][j], err = strconv.Atoi(num)
			if err != nil {
				log.Fatal("Error converting string to int: ", err)
			}
		}
	}

	return sensors
}


func PredictEnd(sensor []int) int {
	var zeroes int = 0
	var prediction int
	var tails []int
	var currentRow []int = sensor

	for zeroes != 1 {
		zeroes = 1
		tails = append(tails, currentRow[len(currentRow) - 1])
		row := make([]int, len(currentRow) - 1)

		for i := range(len(row)) {
			diff := currentRow[i + 1] - currentRow[i]
			if diff != 0 {
				zeroes = 0
			}
			row[i] = diff
		}

		currentRow = row
	}

	for _, tail := range(tails) {
		prediction += tail
	}

	return prediction
}


func Part1() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	sensors := Parse(string(f))

	var sum int

	for _, sensor := range(sensors) {
		sum += PredictEnd(sensor)
	}

	fmt.Println(sum)
}
