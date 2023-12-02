package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
)


func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


func Check2(line string) int {
	var id int
	var body string
	var red, green, blue int

	_, err := fmt.Sscanf(line, "Game %d:", &id)
    if err != nil {
        log.Fatal(err)
    }
	body = strings.TrimPrefix(line, fmt.Sprintf("Game %d:", id))

	draws := strings.Split(body, "; ")

	for _, draw := range draws {

		colors := strings.Split(draw, ", ")

		for _, color := range colors {
			var num int
			var c string

			_, err := fmt.Sscanf(color, "%d %s", &num, &c)
			if err != nil {
				log.Fatal(err)
			}

			if c == "red" {
				red = max(num, red)
			} else if c == "blue" {
				blue = max(num, blue)
			} else if c == "green" {
				green = max(num, green)
			}
		}

	}

	return red * green * blue
}


func Part2() {
    f, err := os.Open("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

	sum := 0

    for scanner.Scan() {
		line := scanner.Text()
		num := Check2(line)
		sum += num
    }

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}
