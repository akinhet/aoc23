package main

import (
	"os"
	"log"
	"fmt"
	"strings"
)


func Parse(f string) ([][]string, int, int) {
	var board [][]string
	var startx, starty int
	lines := strings.Split(f, "\n")
	lines = lines[:len(lines) - 1]

	for i, line := range(lines) {
		board = append(board, strings.Split(line, ""))
		if strings.Contains(line, "S") {
			startx = strings.Index(line, "S")
			starty = i
		}
	}

	return board, startx, starty
}


func ClearBoard(board [][]string) [][]string {
	var clear bool = false
	
	for clear == false {
		clear = true
		for i := range(board) {
			for j := range(board[i]) {
				switch board[i][j] {
				case "|":
					if i == 0 || i == len(board) - 1 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("|F7S", board[i-1][j]) || !strings.Contains("|SLJ", board[i+1][j]){
						board[i][j] = "."
						clear = false
					}
				case "-":
					if j == 0 || j == len(board[0]) - 1 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("S-LF", board[i][j-1]) || !strings.Contains("S-J7", board[i][j+1]){
						board[i][j] = "."
						clear = false
					}
				case "L":
					if i == 0 || j == len(board[0]) - 1 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("|F7S", board[i-1][j]) || !strings.Contains("S-J7", board[i][j+1]){
						board[i][j] = "."
						clear = false
					}
				case "J":
					if i == 0 || j == 0 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("|F7S", board[i-1][j]) || !strings.Contains("S-LF", board[i][j-1]){
						board[i][j] = "."
						clear = false
					}
				case "7":
					if i == len(board) - 1 || j == 0 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("|SLJ", board[i+1][j]) || !strings.Contains("S-LF", board[i][j-1]){
						board[i][j] = "."
						clear = false
					}
				case "F":
					if i == len(board) - 1 || j == len(board[0]) - 1 {
						board[i][j] = "."
						clear = false
					} else if !strings.Contains("|SLJ", board[i+1][j]) || !strings.Contains("S-J7", board[i][j+1]){
						board[i][j] = "."
						clear = false
					}
				}
			}
		}
	}

	return board
}


func CountSteps(board [][]string, startx, starty int) int {
	var x, y int
	var steps int
	var previous string
	
	if strings.Contains("|LJ", board[starty+1][startx]) {
		x = startx
		y = starty+1
		previous = "N"
	} else if strings.Contains("|F7", board[starty-1][startx]) {
		x = startx
		y = starty-1
		previous = "S"
	} else if strings.Contains("-LF", board[starty][startx-1]) {
		x = startx-1
		y = starty
		previous = "E"
	} else if strings.Contains("-J7", board[starty][startx+1]) {
		x = startx+1
		y = starty
		previous = "W"
	}

	for board[y][x] != "S" {
		steps++
		switch board[y][x] {
		case "|":
			if previous == "N" {
				y++
			} else {
				y--
			}
		case "-":
			if previous == "W" {
				x++
			} else {
				x--
			}
		case "L":
			if previous == "N" {
				x++
				previous = "W"
			} else {
				y--
				previous = "S"
			}
		case "J":
			if previous == "N" {
				x--
				previous = "E"
			} else {
				y--
				previous = "S"
			}
		case "7":
			if previous == "S" {
				x--
				previous = "E"
			} else {
				y++
				previous = "N"
			}
		case "F":
			if previous == "S" {
				x++
				previous = "W"
			} else {
				y++
				previous = "N"
			}
		case "S":
			break
		case ".":
			log.Fatal("Went off the path")
		}
	}

	return (steps + 1) / 2
}


func Part1() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput1")
	//f, err := os.ReadFile("testinput2")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	board, startx, starty := Parse(string(f))

	board = ClearBoard(board)
	fmt.Println(CountSteps(board, startx, starty))
	//for _, row := range(board) {
		//for _, c := range(row) {
			//fmt.Print(c)
		//}
		//fmt.Println()
	//}
}
