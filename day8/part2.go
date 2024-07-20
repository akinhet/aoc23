package main

import (
	"fmt"
	"os"
	"log"
	//"github.com/k0kubun/pp"
)


func Move2(instructions string, graph map[string]Node) int {
	var idx int
	var steps int
	var nodes []Node
	
	for name := range(graph) {
		if name[2] == 'A' {
			nodes = append(nodes, graph[name])
		}
	}

	//fmt.Println(nodes)

	for {
		char := instructions[idx]
		solved := true
		//fmt.Println(string(char), nodes)

		for i := range(nodes) {
			if char == 'L' {
				nodes[i] = graph[nodes[i].left]
			} else {
				nodes[i] = graph[nodes[i].right]
			}
			
			if nodes[i].name[2] != 'Z' {
				solved = false
			}
		}

		steps++

		if solved {
			return steps
		}

		idx = (idx + 1) % len(instructions)
	}
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

	fmt.Println(Move2(instructions, graph))
}
