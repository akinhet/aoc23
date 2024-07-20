package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	//"github.com/k0kubun/pp"
)

type Node struct {
	name string
	left, right string
}


func Parse(f string) (string, map[string]Node) {
	var instructions string
	graph := make(map[string]Node)
	lines := strings.Split(f, "\n")

	instructions = lines[0]
	lines = lines[2:len(lines)-1]

	for _, line := range(lines) {
		var n Node
		r := strings.NewReader(line)
		fmt.Println(line)
		num, err := fmt.Fscanf(r, "%3s = (%3s, %3s)", &(n.name), &(n.left), &(n.right))
		if err != nil {
			log.Fatal("Fscanf: ", err, ": ", line)
		} else if num != 3 {
			log.Fatal("Fscanf: not all fields were parsed: ", num)
		}

		graph[n.name] = n
	}

	return instructions, graph
}


func Move(instructions string, graph map[string]Node) int {
	var idx int
	var steps int
	curNode := graph["AAA"]

	for {
		char := instructions[idx]
		if char == 'L' {
			curNode = graph[curNode.left]
		} else {
			curNode = graph[curNode.right]
		}

		steps++

		if curNode.name == "ZZZ" {
			return steps
		}

		idx = (idx + 1) % len(instructions)
	}
}


func Part1() {
	f, err := os.ReadFile("input")
    //f, err := os.ReadFile("testinput1")
    //f, err := os.ReadFile("testinput2")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	instructions, graph := Parse(string(f))

	fmt.Println(instructions)
	fmt.Println(graph)

	fmt.Println(Move(instructions, graph))
}
