package main

import (
	"flag"
	"fmt"

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

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		nrWin := 0
		winning, has := parseLine(line)
		for _, h := range has {
			for _, w := range winning {
				if h == w {
					nrWin++
				}
			}
		}
		if nrWin > 0 {
			win := 1 << (nrWin - 1)
			sum += win
		}
	}

	return sum
}

func task2(lines []string) int {
	count := make(map[int]int)
	for i, line := range lines {
		nrWin := 0
		nr := i + 1
		count[nr] = count[nr] + 1
		winning, has := parseLine(line)
		for _, h := range has {
			for _, w := range winning {
				if h == w {
					nrWin++
				}
			}
		}
		if nrWin > 0 {
			for j := nr + 1; j < nr+1+nrWin; j++ {
				count[j] = count[j] + count[nr]
			}
		}
	}
	sum := 0
	for _, v := range count {
		sum += v
	}

	return sum
}

func parseLine(line string) (winning []int, has []int) {
	ps := u.SplitWithTrim(line, ":")
	ps = u.SplitWithTrim(ps[1], "|")
	return u.SplitToInts(ps[0]), u.SplitToInts(ps[1])
}
