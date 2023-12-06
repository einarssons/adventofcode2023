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

type race struct {
	time, dist int
}

func task1(lines []string) int {
	times := u.SplitToInts(lines[0])
	dists := u.SplitToInts(lines[1])
	p := 1
	for i := 0; i < len(times); i++ {
		p *= nrWins(race{times[i], dists[i]})
	}
	return p
}

func task2(lines []string) int {
	times := u.SplitToInts(lines[0])
	dists := u.SplitToInts(lines[1])
	t := ""
	d := ""
	for i := 0; i < len(times); i++ {
		t += fmt.Sprintf("%d", times[i])
		d += fmt.Sprintf("%d", dists[i])
	}

	r := race{u.Atoi(t), u.Atoi(d)}
	return nrWins(r)
}

func nrWins(r race) int {
	wins := 0
	for i := 1; i < r.time; i++ {
		d := (r.time - i) * i
		if d > r.dist {
			wins++
		}
	}
	return wins
}
