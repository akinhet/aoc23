package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"math"
	"strings"
)


func Check1(line string) int {
	var id int
	var body string

	_, err := fmt.Sscanf(line, "Card %d:", &id)
    if err != nil {
        log.Fatal(err)
    }
	body = strings.TrimPrefix(line, fmt.Sprintf("Card %d:", id))

	nums := strings.Split(body, "|")
	winning := make(map[string]struct{})

	for _, i := range strings.Split(nums[0], " ") {
		if i == "" {
			continue
		}
		winning[i] = struct{}{}
	}

	var sum int

	for _, n := range strings.Split(nums[1], " ") {
		if _, isPresent := winning[n]; isPresent {
			sum += 1
		}
	}

	return sum
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
		sum += int(math.Pow(2, float64(Check1(line) - 1)))
    }

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}
