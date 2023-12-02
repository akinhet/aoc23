package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	sum := 0

	for scanner.Scan() {
		first := -1
		last  := 0

		for _, c := range scanner.Text() {
			if c >= 48 && c <= 57 {
				if first == -1 {
					first = int(c - 48)
				}
				last = int(c - 48)
			}
		}

		sum += first * 10 + last
	}

	fmt.Println(sum)
}
