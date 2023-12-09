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
		gens := process(line)
		val := extendRight(gens)
		sum += val
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for _, line := range lines {
		gens := process(line)
		val := extendLeft(gens)
		sum += val
	}
	return sum
}

func process(line string) [][]int {
	var gens [][]int
	g := u.SplitToInts(line)
	gens = append(gens, g)
	level := 0
topLoop:
	for {
		var allZero bool
		n := make([]int, 0, len(gens[level]))
		for i := 0; i < len(gens[level])-1; i++ {
			allZero = true
			d := gens[level][i+1] - gens[level][i]
			if d != 0 {
				allZero = false
			}
			n = append(n, gens[level][i+1]-gens[level][i])
		}
		gens = append(gens, n)
		level++
		if allZero {
			break topLoop
		}
	}
	return gens
}

func extendRight(gens [][]int) int {
	nrLevels := len(gens)
	gens[nrLevels-1] = append(gens[nrLevels-1], 0)
	for level := len(gens) - 2; level >= 0; level-- {
		last := gens[level][len(gens[level])-1]
		lowerLast := gens[level+1][len(gens[level+1])-1]
		newLast := last + lowerLast
		gens[level] = append(gens[level], newLast)
	}
	return gens[0][len(gens[0])-1]
}

func extendLeft(gens [][]int) int {
	nrLevels := len(gens)
	gens[nrLevels-1] = append([]int{0}, gens[nrLevels-1]...)
	for level := len(gens) - 2; level >= 0; level-- {
		last := gens[level][0]
		lowerLast := gens[level+1][0]
		newLast := last - lowerLast
		gens[level] = append([]int{newLast}, gens[level]...)
	}
	return gens[0][0]
}
