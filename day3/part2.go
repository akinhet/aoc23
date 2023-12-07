package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"log"
)

type Vertex struct {
	Y, X int
}

var m = map[Vertex][]int {}


func Check2(line, prev, next string) map[Vertex][]int {
	var ret = map[Vertex][]int {}

	for i := 0; i < len(line); i++ {
		var startidx, endidx int

		if line[i] < 48 || line[i] > 57 {
			continue
		}

		startidx = i

		for i < len(line) && line[i] >= 48 && line[i] <= 57 {
			endidx = i
			i++
		}

		num, err := strconv.Atoi(line[startidx:endidx+1])
		if err != nil {
			log.Fatal("error converting string to number: ", err)
		}

		
		if startidx - 1 >= 0 {
			if line[startidx - 1] == '*' {
				ret[Vertex{0,  startidx - 1}] = append(ret[Vertex{0,  startidx - 1}], num)
			}
			if prev[startidx - 1] == '*' {
				ret[Vertex{-1, startidx - 1}] = append(ret[Vertex{-1, startidx - 1}], num)
			}
			if next[startidx - 1] == '*' {
				ret[Vertex{1,  startidx - 1}] = append(ret[Vertex{1,  startidx - 1}], num)
			}
		}
		if endidx + 1 < len(line) {
			if line[endidx + 1] == '*' {
				ret[Vertex{0,  endidx + 1}] = append(ret[Vertex{0,  endidx + 1}], num)
			}
			if prev[endidx + 1] == '*' {
				ret[Vertex{-1, endidx + 1}] = append(ret[Vertex{-1, endidx + 1}], num)
			}
			if next[endidx + 1] == '*' {
				ret[Vertex{1,  endidx + 1}] = append(ret[Vertex{1,  endidx + 1}], num)
			}
		}
		for j := startidx; j <= endidx; j++ {
			if prev[j] == '*' {
				ret[Vertex{-1, j}] = append(ret[Vertex{-1, j}], num)
			}
			if next[j] == '*' {
				ret[Vertex{1,  j}] = append(ret[Vertex{1,  j}], num)
			}
		}

	}

	return ret
}


func Part2() {
    //f, err := os.Open("testinput")
    f, err := os.Open("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
	


	var prev, curr, next string
	curr = "............................................................................................................................................"
	scanner.Scan()
	next = scanner.Text()
	var l int

	for scanner.Scan() {
		prev = curr
		curr = next
		next = scanner.Text()

		d := Check2(curr, prev, next)
		for k, v := range d {
			k.Y += l
			//fmt.Println(d, k, v)
			for _, i := range v {
				m[k] = append(m[k], i)
			}
		}

		l++
    }

	prev = curr
	curr = next
	next = "............................................................................................................................................"

	d := Check2(curr, prev, next)

	for k, v := range d {
		k.Y += l
		//fmt.Println(d, k, v)
		for _, i := range v {
			m[k] = append(m[k], i)
		}
	}

	sum := 0
	//fmt.Println(m)

	for _, v := range m {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}

    if err := scanner.Err(); err != nil {
		log.Fatal("error scanning file: ", err)
    }

	fmt.Println(sum)
}

