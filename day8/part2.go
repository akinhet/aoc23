package main

import (
	"fmt"
	"os"
	"log"
	//"github.com/k0kubun/pp"
)


func GCD(a, b int) int {
	if a < b {
		a,b = b,a
	}

	for b > 0 {
		a,b = b,(a % b)
	}

	return a
}


func Move2(instructions string, graph map[string]Node, start string) int {
	var idx int
	var steps int
	curNode := graph[start]

	for {
		char := instructions[idx]
		if char == 'L' {
			curNode = graph[curNode.left]
		} else {
			curNode = graph[curNode.right]
		}

		steps++

		if curNode.name[2] == 'Z' {
			return steps
		}

		idx = (idx + 1) % len(instructions)
	}
}


func CalcSteps(instructions string, graph map[string]Node) int {
	var minSteps int = 1
	
	for name := range(graph) {
		if name[2] == 'A' {
			steps := Move2(instructions, graph, name)
			fmt.Println(steps)
			minSteps = minSteps / GCD(minSteps, steps) * steps
		}
	}

	return minSteps
}


func Part2() {
	f, err := os.ReadFile("input")
    //f, err := os.ReadFile("testinput1")
	//f, err := os.ReadFile("testinput2")
	//f, err := os.ReadFile("testinput3")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	instructions, graph := Parse(string(f))

	fmt.Println(CalcSteps(instructions, graph))
}
