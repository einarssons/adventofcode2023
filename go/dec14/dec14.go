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
		fmt.Println("task2: ", task2(lines, 200, 1_000_000_000))
	}
}

func task1(lines []string) int {
	g := u.CreateRuneGridFromLines(lines)
	tiltNorth(g)
	return countLoad(g)
}

func task2(lines []string, cycleStart, targetLoops int) int {
	g := u.CreateRuneGridFromLines(lines)
	for i := 0; i < cycleStart; i++ {
		tiltNorth(g)
		tiltWest(g)
		tiltSouth(g)
		tiltEast(g)
	}
	sequence := make([]int, 0, cycleStart)
	for nr := 0; nr < cycleStart; nr++ {
		sequence = append(sequence, countLoad(g))
		tiltNorth(g)
		tiltWest(g)
		tiltSouth(g)
		tiltEast(g)
	}
	loop := findCycle(sequence)

	idx := (targetLoops - cycleStart) % len(loop)
	return loop[idx]
}

func findCycle(sequence []int) []int {
	for start := 0; start < len(sequence); start++ {
		for length := 3; length < len(sequence); length++ {
			if isCycle(sequence, start, length) {
				return sequence[start : start+length]
			}
		}
	}
	panic("no cycle found")
}

func isCycle(sequence []int, start, length int) bool {
	for i := 0; i < len(sequence)-length; i++ {
		if sequence[i] != sequence[(i+start)%length] {
			return false
		}
	}
	return true
}

func tiltNorth(g u.RuneGrid) {
	for col := 0; col < g.Width; col++ {
		for {
			change := false
			for row := 1; row < g.Height; row++ {

				if g.Grid[row][col] == 'O' && g.Grid[row-1][col] == '.' {
					g.Grid[row][col] = '.'
					g.Grid[row-1][col] = 'O'
					change = true
				}
			}
			if !change {
				break
			}
		}
	}
}

func tiltSouth(g u.RuneGrid) {
	for col := 0; col < g.Width; col++ {
		for {
			change := false
			for row := g.Height - 2; row >= 0; row-- {
				if g.Grid[row][col] == 'O' && g.Grid[row+1][col] == '.' {
					g.Grid[row][col] = '.'
					g.Grid[row+1][col] = 'O'
					change = true
				}
			}
			if !change {
				break
			}
		}
	}
}

func tiltWest(g u.RuneGrid) {
	for row := 0; row < g.Height; row++ {
		for {
			change := false
			for col := 1; col < g.Width; col++ {
				if g.Grid[row][col] == 'O' && g.Grid[row][col-1] == '.' {
					g.Grid[row][col] = '.'
					g.Grid[row][col-1] = 'O'
					change = true
				}
			}
			if !change {
				break
			}
		}
	}
}

func tiltEast(g u.RuneGrid) {
	for row := 0; row < g.Height; row++ {
		for {
			change := false
			for col := g.Width - 2; col >= 0; col-- {
				if g.Grid[row][col] == 'O' && g.Grid[row][col+1] == '.' {
					g.Grid[row][col] = '.'
					g.Grid[row][col+1] = 'O'
					change = true
				}
			}
			if !change {
				break
			}
		}
	}
}

func countLoad(g u.RuneGrid) int {
	count := 0
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if g.Grid[row][col] == 'O' {
				count += g.Height - row
			}
		}
	}
	return count
}
