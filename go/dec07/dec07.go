package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

type entry struct {
	hand hand
	bid  int
}

type hand string

func (h hand) Type() int {
	bins := make(map[rune]int)
	for _, c := range h {
		bins[c]++
	}
	switch {
	case len(bins) == 1:
		return 6 // Full house
	case len(bins) == 5:
		return 0 // High card
	case len(bins) == 4:
		return 1 // One pair
	case len(bins) == 3:
		for _, v := range bins {
			if v == 2 {
				return 2 // Two pair
			}
		}
		return 3 // Three of a kind
	case len(bins) == 2:
		for _, v := range bins {
			if v == 3 {
				return 4 // Full house
			}
		}
		return 5 // Four of a kind
	default:
		panic("Invalid hand")
	}
}

func compare(a, b hand) bool {
	if a.Type() > b.Type() {
		return true
	} else if a.Type() < b.Type() {
		return false
	}
	for i := 0; i < len(a); i++ {
		av := value(rune(a[i]))
		bv := value(rune(b[i]))
		if av == bv {
			continue
		}
		if av > bv {
			return true
		} else if av < bv {
			return false
		}
	}
	return true
}

func value(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func task1(lines []string) int {
	entries := make([]entry, 0, len(lines))
	for _, line := range lines {
		hand, bid := parse(line)
		entries = append(entries, entry{hand, bid})
	}
	sort.Slice(entries, func(i, j int) bool {
		return compare(entries[i].hand, entries[j].hand)
	})
	sum := 0
	nrEntries := len(entries)
	for i, e := range entries {
		rank := nrEntries - i
		sum += rank * e.bid
	}
	return sum
}

func task2(lines []string) int {
	entries := make([]entry, 0, len(lines))
	for _, line := range lines {
		hand, bid := parse(line)
		entries = append(entries, entry{hand, bid})
	}
	sort.Slice(entries, func(i, j int) bool {
		return compare2(entries[i].hand, entries[j].hand)
	})
	sum := 0
	nrEntries := len(entries)
	for i, e := range entries {
		rank := nrEntries - i
		sum += rank * e.bid
	}
	return sum
}

func parse(line string) (hand, int) {
	parts := strings.Fields(line)
	hand := hand(parts[0])
	bid := u.Atoi(parts[1])
	return hand, bid
}

func (h hand) type2() int {
	bins := make(map[rune]int)
	for _, c := range h {
		bins[c]++
	}
	nrJokers := bins['J']
	switch len(bins) {
	case 1:
		return 6 // Full house
	case 2:
		switch nrJokers {
		case 0:
			for _, v := range bins {
				if v == 3 {
					return 4 // Full house
				}
			}
			return 5 // Four of a kind
		case 1:
			return 6 // Five of a kind (joker + four of a kind)
		case 2:
			return 6 // Five of a kind (2*joker + three of a kind)
		case 3:
			return 6 // Five of a kind (3*joker + pair)
		case 4:
			return 6 // Five of a kind (4*joker + card)
		default:
			panic("Invalid hand 2")
		}
	case 3:
		switch nrJokers {
		case 0:
			for _, v := range bins {
				if v == 2 {
					return 2 // Two pair
				}
			}
			return 3 // Three of a kind
		case 1:
			for _, v := range bins {
				if v == 2 {
					return 4 // Full house (pair + joker + pair)
				}
			}
			return 5 // Four of a kind (3 + joker)
		case 2:
			return 5 // Four of a kind (2*joker + pair)
		case 3:
			return 5 // Four of a kind (3*joker + card)
		default:
			panic("Invalid hand 3")
		}
	case 4:
		switch nrJokers {
		case 0:
			return 1 // One pair
		case 1:
			return 3 // Three of a kind (pair +joker)
		case 2:
			return 3 // Three of a kind (2*joker + 1 card)
		default:
			panic("Invalid hand 4")
		}
	case 5:
		if nrJokers == 1 {
			return 1 // One pair
		}
		return 0 // High card
	default:
		panic("Invalid hand 5")
	}
}

func compare2(a, b hand) bool {
	if a.type2() > b.type2() {
		return true
	} else if a.type2() < b.type2() {
		return false
	}
	for i := 0; i < len(a); i++ {
		av := value2(rune(a[i]))
		bv := value2(rune(b[i]))
		if av == bv {
			continue
		}
		if av > bv {
			return true
		} else if av < bv {
			return false
		}
	}
	return true
}

func value2(card rune) int {
	switch card {

	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J': // Joker
		return 1
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}
