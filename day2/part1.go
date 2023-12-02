package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
)


func Check1(line string) int {
	var id int
	var body string

	_, err := fmt.Sscanf(line, "Game %d:", &id)
    if err != nil {
        log.Fatal(err)
    }
	body = strings.TrimPrefix(line, fmt.Sprintf("Game %d:", id))

	draws := strings.Split(body, "; ")

	for _, draw := range draws {
		var red, green, blue int

		colors := strings.Split(draw, ", ")

		for _, color := range colors {
			var num int
			var c string

			_, err := fmt.Sscanf(color, "%d %s", &num, &c)
			if err != nil {
				log.Fatal(err)
			}

			if c == "red" {
				red = num
			} else if c == "blue" {
				blue = num
			} else if c == "green" {
				green = num
			}
		}

		if red > 12 || green > 13 || blue > 14 {
			return 0
		}
	}

	return id
}


func Part1() {
    f, err := os.Open("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

	sum := 0

    for scanner.Scan() {
		line := scanner.Text()
		num := Check1(line)
		sum += num
    }

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}
