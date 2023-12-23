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
	g := u.CreateRuneGridFromLines(lines)
	fmt.Println("width: ", g.Width, "height: ", g.Height)
	startPoint := u.Pos2D{0, 1}
	endPoint := u.Pos2D{g.Height - 1, g.Width - 2}
	g.Grid[startPoint.Row][startPoint.Col] = 'x'
	nrSteps := 0
	pos := startPoint.Add(down)
	cost := longestPath(g, nrSteps, pos, endPoint)
	return cost
}

func task2(lines []string) int {
	g := u.CreateRuneGridFromLines(lines)
	fmt.Println("width: ", g.Width, "height: ", g.Height)
	startPoint := u.Pos2D{0, 1}
	endPoint := u.Pos2D{g.Height - 1, g.Width - 2}
	g.Grid[startPoint.Row][startPoint.Col] = 'x'
	nrSteps := 0
	pos := startPoint.Add(down)
	cost := longestPath2(g, nrSteps, pos, endPoint)
	return cost
}

var (
	up    = u.Pos2D{-1, 0}
	down  = u.Pos2D{1, 0}
	left  = u.Pos2D{0, -1}
	right = u.Pos2D{0, 1}
	dirs  = []u.Pos2D{up, down, left, right}
)

func longestPath(g u.RuneGrid, nrSteps int, start, end u.Pos2D) int {
	g.Grid[start.Row][start.Col] = 'X'
	//fmt.Println(g)
	g.Grid[start.Row][start.Col] = 'x'
	nrSteps++
	if start == end {
		return nrSteps
	}
	ways := make([]u.Pos2D, 0, 2)
	for _, dir := range dirs {
		newPos := start.Add(dir)
		c := g.Grid[newPos.Row][newPos.Col]
		switch c {
		case 'x', '#':
			continue
		case '.':
			ways = append(ways, newPos)
		case '>':
			if dir == right {
				ways = append(ways, newPos)
			}
		case '<':
			if dir == left {
				ways = append(ways, newPos)
			}
		case '^':
			if dir == up {
				ways = append(ways, newPos)
			}
		case 'v':
			if dir == down {
				ways = append(ways, newPos)
			}
		default:
			fmt.Printf("unknown char: '%c'\n", c)
			fmt.Println(g)
			panic("unknown char")

		}
	}
	if len(ways) == 0 {
		return -1
	}
	longest := -1
	for _, way := range ways {
		ng := g
		if len(ways) > 1 {
			ng = g.Copy()
		}
		wayLen := longestPath(ng, nrSteps, way, end)
		if wayLen > longest {
			addLongest(wayLen)
			longest = wayLen
		}
	}
	return longest
}

func longestPath2(g u.RuneGrid, nrSteps int, start, end u.Pos2D) int {
	g.Grid[start.Row][start.Col] = 'X'
	//fmt.Println(g)
	g.Grid[start.Row][start.Col] = 'x'
	nrSteps++
	if start == end {
		return nrSteps
	}
	ways := make([]u.Pos2D, 0, 2)
	for _, dir := range dirs {
		newPos := start.Add(dir)
		c := g.Grid[newPos.Row][newPos.Col]
		switch c {
		case 'x', '#':
			continue
		case '.', '>', '<', '^', 'v':
			ways = append(ways, newPos)
		default:
			fmt.Printf("unknown char: '%c'\n", c)
			fmt.Println(g)
			panic("unknown char")

		}
	}
	if len(ways) == 0 {
		return -1
	}
	longest := -1
	for _, way := range ways {
		ng := g
		if len(ways) > 1 {
			ng = g.Copy()
		}
		wayLen := longestPath2(ng, nrSteps, way, end)
		if wayLen > longest {
			addLongest(wayLen)
			longest = wayLen
		}
	}
	return longest
}

var longest = -1

func addLongest(wayLen int) {
	if wayLen > longest {
		fmt.Printf("found end after %d steps\n", wayLen)
		longest = wayLen
	}
}
