#!/bin/bash

mkdir $1
cd $1
go mod init $1

cat << EOF > main.go
package main

func main() {
	Part1()
	//Part2()
}
EOF

cat << EOF > part1.go
package main

import (
	"os"
	"log"
	"strings"
)


func Parse(f string) {
	lines := strings.Split(f, "\n")
	for _, line := range(lines) {
		
	}
}


func Part1() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	Parse(string(f))
}
EOF

cat << EOF > part2.go
package main

import (
	"os"
	"log"
)


func Part2() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	Parse(string(f))
}
EOF
