package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
)

var nums = map[string]int{
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}


func Extract(line string) int {
	var first, last int
	var firstidx, lastidx int = -1, -1

	for idx, char := range line {
		if char >= 48 && char <= 57 {
			if firstidx == -1 {
				first = int(char - 48)
				firstidx = idx
			}

			last = int(char - 48)
			lastidx = idx
		}
	}

	for str, num := range nums {
		idx := strings.Index(string(line), str)
		if idx == -1 {
			continue
		}

		if idx < firstidx || firstidx == -1 {
			first = num
			firstidx = idx
		}

		idx = strings.LastIndex(string(line), str)
		if idx == -1 {
			continue
		}
		
		if idx > lastidx || lastidx == -1 {
			last = num
			lastidx = idx
		}
	}

	return first * 10 + last
}


func Part2() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	sum := 0

	for scanner.Scan() {
		line := string(scanner.Text())
		
		num := Extract(line)
		sum += num

		fmt.Printf("%s: %d\n", line, num)
	}

	fmt.Println(sum)
}
