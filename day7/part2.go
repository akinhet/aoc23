package main

import (
	"os"
	"log"
	"strings"
	"strconv"
	"github.com/k0kubun/pp"
	"slices"
)


func Parse2(f string) []Hand {
	var ret []Hand

	lines := strings.Split(f, "\n")

	for _, line := range(lines) {
		var hand Hand
		l := strings.Split(line, " ")

		if len(l) < 2 {
			break
		}

		hand.handstr = l[0]

		for _, e := range(l[0]) {
			if s, err := strconv.Atoi(string(e)); err == nil {
				hand.hand = append(hand.hand, byte(s))
			} else {
				var num int

				switch e {
				case 'T':
					num = 10
				case 'J':
					num = 1
				case 'Q':
					num = 12
				case 'K':
					num = 13
				case 'A':
					num = 14
				}

				hand.hand = append(hand.hand, byte(num))
			}
		}
		if s, err := strconv.Atoi(l[1]); err == nil {
			hand.bid = s
		}

		ret = append(ret, hand)
	}

	return ret
}


func CalcHandType2(h Hand) int {
	var hand string = h.handstr
	
	// Five of a kind
	if strings.Count(hand, string(hand[0])) == 5 {
		return 7
	} else if temp := strings.ReplaceAll(hand, "J", ""); strings.Count(temp, string(temp[0])) == len(temp) {
		return 7
	}

	// Four of a kind
	for _, i := range(hand) {
		if strings.Count(hand, string(i)) == 4 {
			return 6
		} else {
			temp := strings.ReplaceAll(hand, "J", "")
			temp = strings.ReplaceAll(temp, string(i), "")
			if len(temp) == 1 {
				return 6
			}
		}
	}

	// Full house
	for _, i := range(hand) {
		if strings.Count(hand, string(i)) == 3 {
			temp := strings.ReplaceAll(hand, string(i), "")
			if temp[0] == temp[1] {
				return 5
			}
		} else if strings.Count(hand, string(i)) == 2 {
			temp := strings.ReplaceAll(hand, string(i), "")
			temp = strings.ReplaceAll(temp, "J", "")
			if strings.Count(temp, string(temp[0])) == len(temp) {
				return 5
			}
		}
	}

	// Three of a kind
	for _, i := range(hand) {
		if strings.Count(hand, string(i)) == 3 {
			return 4
		} else if temp := strings.ReplaceAll(hand, "J", ""); strings.Count(temp, string(i)) == len(temp) - 2 {
			return 4
		}
	}

	// Two pair
	for _, i := range(hand) {
		if strings.Count(hand, string(i)) == 2 {
			temp := strings.ReplaceAll(hand, string(i), "")
			if temp[0] == temp[1] || temp[0] == temp[2] || temp[1] == temp[2] {
				return 3
			} else if strings.Count(temp, "J") == 1 {
				return 3
			}
		}
	}

	// One pair
	for _, i := range(hand) {
		if strings.Count(hand, string(i)) == 2 {
			return 2
		} else if strings.Count(hand, "J") >= 1 {
			return 2
		}
	}

	// High card
	return 1
}


func Part2() {
    f, err := os.ReadFile("input")
    if err != nil {
		log.Fatal("error opening file: ", err)
    }

	hands := Parse2(string(f))
	
	for i := range(hands) {
		hands[i].handtype = CalcHandType2(hands[i])
	}

	slices.SortStableFunc(hands, func(a, b Hand) int {
		if a.handtype > b.handtype {
			return 1
		} else if a.handtype < b.handtype {
			return -1
		}

		for i := range(a.hand) {
			if a.hand[i] > b.hand[i] {
				return 1
			} else if a.hand[i] < b.hand[i] {
				return -1
			}
		}

		return 0
	})
	

	//slices.Reverse(hands)

	var winnings int

	for i, e := range(hands) {
		winnings += e.bid * (i + 1)
		pp.Println(e.handstr, e.bid, e.handtype)
	}
	//pp.Println(hands)

	//var a Hand
	//a.handstr = "5JJJJ"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "53JJJ"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "22J33"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "22J33"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12J33"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12J33"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12JJ8"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12J28"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12J38"
	//pp.Println(a.handstr, CalcHandType2(a))
	//a.handstr = "12438"
	//pp.Println(a.handstr, CalcHandType2(a))

	pp.Println(winnings)
}
