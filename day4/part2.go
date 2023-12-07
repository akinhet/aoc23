package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)


type Card struct {
	wins, count int
}


/*
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
*/


func Part2() {
    f, err := os.Open("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

	cards := []Card{}

	sum := 0

    for scanner.Scan() {
		line := scanner.Text()
		num := Check1(line)
		cards = append(cards, Card{num, 1})
    }

	for i, e := range cards {
		sum += e.count

		for j := 1; j <= e.wins; j++ {
			cards[i + j].count += e.count
		}
	}

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}
