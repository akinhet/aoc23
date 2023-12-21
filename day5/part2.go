package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"github.com/k0kubun/pp"
	"math"
)


func Parse2(file string) Almanac {
	var ret Almanac

	// init scanner
	scanner := bufio.NewScanner(strings.NewReader(file))

	//scan for seeds
	scanner.Scan()
	if text := scanner.Text(); text != "" {
		seedtext, _ := strings.CutPrefix(text, "Seeds: ")
		seeds := strings.Split(seedtext, " ")

		for _, e := range seeds {
			if n, err := strconv.Atoi(e); err == nil {
				ret.seeds = append(ret.seeds, n)
			}
		}
	}
	scanner.Scan()

	// scan for ss
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "seed-to-soil map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.ss = append(ret.ss, m)
	}

	// scan for sf
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "soil-to-fertilizer map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.sf = append(ret.sf, m)
	}

	// scan for fw
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "fertilizer-to-water map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.fw = append(ret.fw, m)
	}

	// scan for wl
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "water-to-light map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.wl = append(ret.wl, m)
	}

	// scan for lt
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "light-to-temperature map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.lt = append(ret.lt, m)
	}

	// scan for th
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "temperature-to-humidity map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.th = append(ret.th, m)
	}

	// scan for hl
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan() // skip the empty line
			break
		}
		if text == "humidity-to-location map:" {
			continue
		}

		nums := strings.Split(text, " ")
		var m Map

		for i, e := range nums {
			if n, err := strconv.Atoi(e); err == nil {
				if i == 0 {
					m.to = n
				} else if i == 1 {
					m.from = n
				} else if i == 2 {
					m.rg = n
				}
			} else {
				log.Println(e, err)
			}
		}

		ret.hl = append(ret.hl, m)
	}

	return ret
}


func ConvertSeedsToLocation2(m Almanac) int {
	var ret int = math.MaxInt

	for i := 0; i < len(m.seeds); i += 2 {
		first := m.seeds[i]
		rng := m.seeds[i+1]

		for j := first; j < first + rng; j++ {
			n := ConvertUsingMap(j, m.ss)
			n = ConvertUsingMap(n, m.sf)
			n = ConvertUsingMap(n, m.fw)
			n = ConvertUsingMap(n, m.wl)
			n = ConvertUsingMap(n, m.lt)
			n = ConvertUsingMap(n, m.th)
			n = ConvertUsingMap(n, m.hl)

			ret = min(ret, n)
		}
	}

	return ret
}


func Part2() {
    f, err := os.ReadFile("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	m := Parse2(string(f))
	seeds := ConvertSeedsToLocation2(m)

	pp.Println(m.seeds)
	//fmt.Println(slices.Min(seeds))
	fmt.Println(seeds)
	//pp.Println(seeds)
	
}
