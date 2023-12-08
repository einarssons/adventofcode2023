package main

import (
	"flag"
	"fmt"
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

	// 15299095376265 to high
}

func task1(lines []string) int {
	directions, mapping := parse(lines)
	pos := "AAA"
	step := 0
	nrSteps := 0
	for {
		if directions[step] == "L" {
			pos = mapping[pos].L
		} else {
			pos = mapping[pos].R
		}
		nrSteps++
		if pos == "ZZZ" {
			return nrSteps
		}
		step++
		step = step % len(directions)
	}
}

type cycle struct {
	offset, period int
}

func task2(lines []string) int {
	directions, mapping := parse(lines)
	aNodes := make([]string, 0)
	for k, _ := range mapping {
		if strings.HasSuffix(k, "A") {
			aNodes = append(aNodes, k)
		}
	}
	fmt.Printf("startNodes: %v\n", aNodes)
	//fmt.Println(directions, aNodes, mapping)
	poss := aNodes
	delta := 0
	cycles := make([]u.Cycle, len(poss))
	for i := range poss {
		step := 0
		nrSteps := 0
		lastNrSteps := 0
		for {
			if directions[step] == "L" {
				poss[i] = mapping[poss[i]].L
			} else {
				poss[i] = mapping[poss[i]].R
			}
			nrSteps++
			if strings.HasSuffix(poss[i], "Z") {
				newDelta := nrSteps - lastNrSteps
				if newDelta == delta {
					cycles[i] = u.Cycle{Offset: nrSteps % newDelta, Period: newDelta}
					break
				}
				delta = nrSteps - lastNrSteps
				lastNrSteps = nrSteps
			}
			step++
			step = step % len(directions)
		}
	}
	fmt.Printf("cycles: %+v\n", cycles)
	return u.CRT(cycles)
}

type direction struct {
	L, R string
}

func parse(lines []string) ([]string, map[string]direction) {
	var directions []string
	mapping := make(map[string]direction)
	for i, line := range lines {
		if i == 0 {
			directions = u.SplitToChars(line)
			continue
		}
		if line == "" {
			continue
		}
		parts := u.SplitWithTrim(line, "=")
		key := parts[0]
		parts = u.SplitWithTrim(parts[1][1:len(parts[1])-1], ",")
		mapping[key] = direction{parts[0], parts[1]}
	}
	return directions, mapping
}
