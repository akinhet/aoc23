package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"math"
)


type point struct {
	x, y int
}


func FloodFill(board [][]string, x, y int) ([][]string, int) {
	queue := make([]point, 0)
	var count int

	queue = append(queue, point{x, y})

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if board[current.y][current.x] != "." {
			continue
		}
		board[current.y][current.x] = "X"

		count++

		if current.x != len(board[0]) - 1 {
			queue = append(queue, point{current.x + 1, current.y})
		}
		if current.x != 0 {
			queue = append(queue, point{current.x - 1, current.y})
		}
		if current.y != len(board) - 1 {
			queue = append(queue, point{current.x, current.y + 1})
		}
		if current.y != 0 {
			queue = append(queue, point{current.x, current.y - 1})
		}
	}


	return board, count
}


func CountInside(board [][]string) int {
	var count, addition int
	var previous string = "N"
	var inside point = point{1, 0}
	var starty, startx int
	var x, y int

	out:
	for i := range(board) {
		for j := range(board[i]) {
			if board[i][j] != "." {
				startx, starty = j, i
				break out
			}
		}
	}

	x, y = startx, starty + 1

	for y != starty || x != startx {
		board, addition = FloodFill(board, x + inside.x, y + inside.y)
		count += addition
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
			inside.x, inside.y = -inside.y, -inside.x
		case "J":
			if previous == "N" {
				x--
				previous = "E"
			} else {
				y--
				previous = "S"
			}
			inside.x, inside.y = inside.y, inside.x
		case "7":
			if previous == "S" {
				x--
				previous = "E"
			} else {
				y++
				previous = "N"
			}
			inside.x, inside.y = -inside.y, -inside.x
		case "F":
			if previous == "S" {
				x++
				previous = "W"
			} else {
				y++
				previous = "N"
			}
			inside.x, inside.y = inside.y, inside.x
		case "S":
			break
		case ".":
			log.Fatal("Went off the path")
		}
	}

	return count
}


func ClearBoard2(board [][]string, startx, starty int) ([][]string, []int, []int, int) {
	var x, y int
	var previous string
	var clearedBoard [][]string = make([][]string, len(board))
	var xs, ys []int
	var boundary int = 1

	for i := range(board) {
		clearedBoard[i] = make([]string, len(board[i]))
		for j := range(clearedBoard[i]) {
			clearedBoard[i][j] = "."
		}
	}

	if starty < len(board) - 1 && startx < len(board[0]) - 1 && strings.Contains("|LJ", board[starty+1][startx]) && strings.Contains("-J7", board[starty][startx+1]) {
		x = startx
		y = starty+1
		previous = "N"
		clearedBoard[starty][startx] = "F"
		xs = append(xs, startx)
		ys = append(ys, starty)
	} else if starty < len(board) - 1 && starty > 0 && strings.Contains("|LJ", board[starty+1][startx]) && strings.Contains("|F7", board[starty-1][startx]) {
		x = startx
		y = starty+1
		previous = "N"
		clearedBoard[starty][startx] = "|"
	} else if starty < len(board) - 1 && startx > 0 && strings.Contains("|LJ", board[starty+1][startx]) && strings.Contains("-LF", board[starty][startx-1]) {
		x = startx
		y = starty+1
		previous = "N"
		clearedBoard[starty][startx] = "7"
		xs = append(xs, startx)
		ys = append(ys, starty)
	} else if starty > 0 && startx > 0 && strings.Contains("|F7", board[starty-1][startx]) && strings.Contains("-LF", board[starty][startx-1]) {
		x = startx
		y = starty-1
		previous = "S"
		clearedBoard[starty][startx] = "J"
		xs = append(xs, startx)
		ys = append(ys, starty)
	} else if starty > 0 && startx < len(board[0]) - 1 && strings.Contains("|F7", board[starty-1][startx]) && strings.Contains("-J7", board[starty][startx+1]) {
		x = startx
		y = starty-1
		previous = "S"
		clearedBoard[starty][startx] = "L"
		xs = append(xs, startx)
		ys = append(ys, starty)
	} else if startx < len(board[0]) - 1 && startx > 0 && strings.Contains("-J7", board[starty][startx+1]) && strings.Contains("-LF", board[starty][startx-1]) {
		x = startx+1
		y = starty
		previous = "W"
		clearedBoard[starty][startx] = "-"
	}

	for y != starty || x != startx {
		boundary++
		clearedBoard[y][x] = board[y][x]
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
			xs = append(xs, x)
			ys = append(ys, y)
			if previous == "N" {
				x++
				previous = "W"
			} else {
				y--
				previous = "S"
			}
		case "J":
			xs = append(xs, x)
			ys = append(ys, y)
			if previous == "N" {
				x--
				previous = "E"
			} else {
				y--
				previous = "S"
			}
		case "7":
			xs = append(xs, x)
			ys = append(ys, y)
			if previous == "S" {
				x--
				previous = "E"
			} else {
				y++
				previous = "N"
			}
		case "F":
			xs = append(xs, x)
			ys = append(ys, y)
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

	xs = append(xs, startx)
	ys = append(ys, starty)

	return clearedBoard, xs, ys, boundary
}


func CalcAreaSholace(xs, ys []int) int {
	var area, idx int

	for idx < (len(xs) - 1) {
		area += xs[idx] * ys[idx + 1]
		area -= xs[idx + 1] * ys[idx]

		idx++
	}

	return int(math.Abs(float64(area) / 2))
}


func Part2() {
	f, err := os.ReadFile("input")
	//f, err := os.ReadFile("testinput1")
	//f, err := os.ReadFile("testinput2")
	//f, err := os.ReadFile("testinput3")
	//f, err := os.ReadFile("testinput4")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	board, startx, starty := Parse(string(f))
	//fmt.Println(startx, starty, board)
	fmt.Println()
	board, xs, ys, boundary := ClearBoard2(board, startx, starty)
	for _, row := range(board) {
		for _, c := range(row) {
			fmt.Print(c)
		}
		fmt.Println()
	}

	//fmt.Println(xs, "\n", ys)
	area := CalcAreaSholace(xs, ys)

	fmt.Println(area + 1 - boundary / 2)
	fmt.Println(CountInside(board))
	//for _, row := range(board) {
		//for _, c := range(row) {
			//fmt.Print(c)
		//}
		//fmt.Println()
	//}
}
