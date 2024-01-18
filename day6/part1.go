package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"github.com/k0kubun/pp"
)


func Parse(f string) ([]int, []int) {
	lines := strings.Split(f, "\n")

	timesunparsed, _ := strings.CutPrefix(lines[0], "Time: ")
	timestext := strings.Split(timesunparsed, " ")

	var times []int

	for _, e := range timestext {
		if n, err := strconv.Atoi(e); err == nil {
			times = append(times, n) 
		}
	}

	recordsunparsed, _ := strings.CutPrefix(lines[1], "Distance: ")
	recordstext := strings.Split(recordsunparsed, " ")

	var records []int

	for _, e := range recordstext {
		if n, err := strconv.Atoi(e); err == nil {
			records = append(records, n) 
		}
	}

	return times, records
}


func WaysToBeatRecord(time, record int) int {
	min := 0
	for min < time {
		distance := (time - min) * min
		
		if distance > record {
			break
		}
		min++
	}

	max := time - min

	pp.Println(time, record, min, max)

	return max - min + 1
}


func Part1() {
    f, err := os.ReadFile("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	times, records := Parse(string(f))

	ways := 1

	for i := range times {
		ways *= WaysToBeatRecord(times[i], records[i])
	}

	fmt.Println(times, records)
	fmt.Println(ways)
}
