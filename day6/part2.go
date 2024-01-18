package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"github.com/k0kubun/pp"
)


func Parse2(f string) (int, int) {
	lines := strings.Split(f, "\n")

	timesunparsed, _ := strings.CutPrefix(lines[0], "Time: ")
	timestext := strings.ReplaceAll(timesunparsed, " ", "")

	var time int

	if n, err := strconv.Atoi(timestext); err == nil {
		time = n
	}

	recordsunparsed, _ := strings.CutPrefix(lines[1], "Distance: ")
	recordstext := strings.ReplaceAll(recordsunparsed, " ", "")

	var record int

	if n, err := strconv.Atoi(recordstext); err == nil {
		record = n
	}

	return time, record
}


func WaysToBeatRecord2(time, record int) int {
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


func Part2() {
    f, err := os.ReadFile("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	time, record := Parse2(string(f))

	ways := WaysToBeatRecord(time, record)

	fmt.Println(ways)
}
