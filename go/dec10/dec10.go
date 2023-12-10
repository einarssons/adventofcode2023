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
	g := u.CreateCharGridFromLines(lines)
	l := findLoop(g)
	loopLen := len(l)
	longestDist := loopLen / 2
	return longestDist
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	l := findLoop(g)
	//fmt.Println("loop", l)
	//fmt.Println(g)
	g2 := u.CreateCharGridFromLines(lines)
	markPath(g2, l)
	//fmt.Println(g2)
	markNonPath(g, g2)
	//fmt.Println("")
	//fmt.Println(g)
	nrEnclosed := countEnclosedAreas(g)
	//fmt.Println("")
	//fmt.Println(g)
	return nrEnclosed
}

func markPath(g u.CharGrid, path []pos) {
	for _, p := range path {
		g.Grid[p.r][p.c] = "X"
	}
}

// Modify g to mark all non-path from g2 as "."
func markNonPath(g, g2 u.CharGrid) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g2.Grid[r][c] != "X" {
				g.Grid[r][c] = "."
			}
		}
	}
}

// countEnclosedAreas returns the number path crossings on one side of a path
func countEnclosedAreas(g u.CharGrid) int {
	count := 0
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == "." {
				nrV := nrVerticalCrossings(g, r, c)
				nrH := nrHorizontalCrossings(g, r, c)
				if nrV%2 == 1 && nrH%2 == 1 {
					g.Grid[r][c] = "I"
					count++
				} else {
					g.Grid[r][c] = "O"
				}
			}
		}
	}
	return count
}

// nrVerticalCrossings returns the number of vertical crossings to one edge of a row
func nrVerticalCrossings(g u.CharGrid, row, col int) int {
	nrCrossings := 0
	if col == 0 || col == g.Width-1 {
		return 0
	}
	startC, endC := 0, col
	for c := 0; c < col; c++ {
		if g.Grid[row][c] == "S" {
			startC, endC = col+1, g.Width
		}
	}

	// levels is a stack with directions from where we entered
	levels := make([]string, 0, 5)
	for c := startC; c < endC; c++ {
		chr := g.Grid[row][c]
		switch chr {
		case "|":
			nrCrossings++
		case "F":
			levels = append(levels, "bottom")
		case "L":
			levels = append(levels, "top")
		case "J":
			l := levels[len(levels)-1]
			if l == "bottom" {
				nrCrossings++
			}
			levels = levels[:len(levels)-1]
		case "7":
			l := levels[len(levels)-1]
			if l == "top" {
				nrCrossings++
			}
			levels = levels[:len(levels)-1]
		case "-":
			if len(levels) == 0 {
				panic("should be inside when seeing -")
			}
		}
	}
	return nrCrossings
}

// nrVerticalCrossings returns the number of horizontal crossings to one edge of a column
func nrHorizontalCrossings(g u.CharGrid, row, col int) int {
	nrCrossings := 0
	if row == 0 || row == g.Height-1 {
		return 0
	}
	startR, endR := 0, row
	for r := 0; r < row; r++ {
		if g.Grid[r][col] == "S" {
			startR, endR = row+1, g.Height
		}
	}
	// levels is a stack with directions from where we entered
	levels := make([]string, 0, 5)
	for r := startR; r < endR; r++ {
		chr := g.Grid[r][col]
		switch chr {
		case "-":
			nrCrossings++
		case "7":
			levels = append(levels, "left")
		case "F":
			levels = append(levels, "right")
		case "J":
			l := levels[len(levels)-1]
			if l == "right" {
				nrCrossings++
			}
			levels = levels[:len(levels)-1]
		case "L":
			l := levels[len(levels)-1]
			if l == "left" {
				nrCrossings++
			}
			levels = levels[:len(levels)-1]
		case "|":
			if len(levels) == 0 {
				panic("should be inside when seeing |")
			}
		}
	}
	return nrCrossings
}

type pos struct {
	r, c int
}

func findLoop(g u.CharGrid) []pos {
	var start pos
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == "S" {
				start = pos{r, c}
				break
			}
		}
	}
	path := startPath(g, start)
	loop := traverse(g, path)
	return loop
}

func startPath(g u.CharGrid, p pos) []pos {
	path := make([]pos, 0, 10000)
	path = append(path, p)
	nps := []pos{{p.r - 1, p.c},
		{p.r + 1, p.c}, {p.r, p.c - 1}, {p.r, p.c + 1}}
	for _, np := range nps {
		if !g.InBounds(np.r, np.c) {
			continue
		}
		c := g.Grid[np.r][np.c]
		ok := false
		switch {
		case c == "|" && (p.r == np.r+1 || p.r == np.r-1):
			ok = true
		case c == "-" && (p.c == np.c+1 || p.c == np.c-1):
			ok = true
		case c == "L" && (p.c == np.c-1 || p.r == np.r+1):
			ok = true
		case c == "J" && (np.c == p.c+1 || np.r == p.r+1):
			ok = true
		case c == "F" && (np.c == p.c-1 || np.r == p.r-1):
			ok = true
		case c == "7" && (np.c == p.c+1 || np.r == p.r-1):
			ok = true
		}
		if ok {
			path = append(path, np)
			break
		}
	}
	return path
}

// Return new path and then looplen when back to start
// The latest element of path is a neighbor of the previous element
// We know that it was reached from the previous element
func traverse(g u.CharGrid, path []pos) []pos {
	for {
		p := path[len(path)-1]
		nps := []pos{{p.r - 1, p.c}, {p.r + 1, p.c}, {p.r, p.c - 1}, {p.r, p.c + 1}}
		prev := path[len(path)-2]
		c := g.Grid[p.r][p.c]
	npsLoop:
		for _, np := range nps {
			if !g.InBounds(np.r, np.c) || np == prev {
				continue
			}
			nc := g.Grid[np.r][np.c]
			if nc == "." {
				continue
			}
			dir := ""
			switch c {
			case "|":
				if p.r == prev.r+1 { // From top
					if np.r == p.r+1 { // OK from top
						dir = "down"
					}
				} else { // From bottom
					if np.r == p.r-1 { // OK from bottom
						dir = "up"
					}
				}
			case "-":
				if p.c == prev.c+1 { // From left
					if np.c == p.c+1 { // OK from left
						dir = "right"
					}
				} else { // From right
					if np.c == p.c-1 { // OK from right
						dir = "left"
					}
				}
			case "L":
				if p.c == prev.c-1 { // From right
					if np.r == p.r-1 { // OK from bottom
						dir = "up"
					}
				} else { // From top
					if np.c == p.c+1 { // OK from left
						dir = "right"
					}
				}
			case "J":
				if p.c == prev.c+1 { // From left
					if np.r == p.r-1 { // OK from bottom
						dir = "up"
					}
				} else { // From top
					if np.c == p.c-1 { // OK from right
						dir = "left"
					}
				}
			case "F":
				if p.c == prev.c-1 { // From right
					if np.r == p.r+1 { // OK from top
						dir = "down"
					}
				} else { // From bottom
					if np.c == p.c+1 { // OK from left
						dir = "right"
					}
				}
			case "7":
				if p.c == prev.c+1 { // From left
					if np.r == p.r+1 { // OK from top
						dir = "down"
					}
				} else { // From bottom
					if np.c == p.c-1 { // OK from right
						dir = "left"
					}
				}
			}
			if dir == "" {
				continue
			}
			if nc == "S" {
				fmt.Printf("found loop of length %d\n", len(path))
				path = append(path, np)
				return path
			}
			goodDir := false
			switch dir {
			case "up":
				switch nc {
				case "|", "F", "7":
					goodDir = true
				}
			case "down":
				switch nc {
				case "|", "L", "J":
					goodDir = true
				}
			case "left":
				switch nc {
				case "-", "F", "L":
					goodDir = true
				}
			case "right":
				switch nc {
				case "-", "7", "J":
					goodDir = true
				}
			}
			if goodDir {
				path = append(path, np)
				break npsLoop
			}
		}
	}
}
