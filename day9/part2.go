package main

import (
	"os"
	"log"
	"fmt"
)


func PredictStart(sensor []int) int {
	var zeroes int = 0
	var prediction int
	var heads []int
	var currentRow []int = sensor

	for zeroes != 1 {
		zeroes = 1
		heads = append(heads, currentRow[0])
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

	for i, head := range(heads) {
		if i % 2 == 1 {
			head *= -1
		}
		prediction += head
	}

	return prediction
}


func Part2() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	sensors := Parse(string(f))

	var sum int

	for _, sensor := range(sensors) {
		sum += PredictStart(sensor)
	}

	fmt.Println(sum)
}
