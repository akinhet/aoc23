package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"log"
)


func Check1(line, prev, next string) int {
	var sum int

	for i := 0; i < 140; i++ {
		var startidx, endidx int
		var isserial bool

		if line[i] < 48 || line[i] > 57 {
			continue
		}

		startidx = i
		
		for i < 140 && line[i] >= 48 && line[i] <= 57 {
			endidx = i
			i++
		}

		num, err := strconv.Atoi(line[startidx:endidx+1])
		if err != nil {
			log.Fatal("error converting string to number: ", err)
		}
		
		if startidx - 1 >= 0 {
			if line[startidx - 1] != '.' || prev[startidx - 1] != '.' || next[startidx - 1] != '.' {
				isserial = true
			}
		}
		if endidx + 1 < 140 {
			if line[endidx + 1] != '.' || prev[endidx + 1] != '.' || next[endidx + 1] != '.' {
				isserial = true
			}
		}
		for j := startidx; j <= endidx; j++ {
			if prev[j] != '.' || next[j] != '.' {
				isserial = true
			}
		}

		if isserial {
			sum += num
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

	var prev, curr, next string
	curr = "............................................................................................................................................"
	scanner.Scan()
	next = scanner.Text()

	sum := 0

    for scanner.Scan() {
		prev = curr
		curr = next
		next = scanner.Text()

		num := Check1(curr, prev, next)
		sum += num
    }

	prev = curr
	curr = next
	next = "............................................................................................................................................"
	sum += Check1(curr, prev, next)

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}
