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
	l := findLongestPath(g)
	return l
}

func task2(lines []string) int {
	return 0
}

type pos struct {
	r, c int
}

func findLongestPath(g u.CharGrid) int {
	var start pos
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == "S" {
				start = pos{r, c}
				break
			}
		}
	}
	fmt.Println("start", start)
	path := startPath(g, start)
	loopLen := traverse(g, path) - 1
	longestDist := loopLen/2 + loopLen%2
	return longestDist
}

func startPath(g u.CharGrid, p pos) []pos {
	var path []pos
	path = append(path, p)
	nps := []pos{pos{p.r - 1, p.c},
		pos{p.r + 1, p.c}, pos{p.r, p.c - 1}, pos{p.r, p.c + 1}}
	for _, np := range nps {
		if !g.InBounds(np.r, np.c) {
			continue
		}
		c := g.Grid[np.r][np.c]
		switch c {
		case "|":
			if p.r == np.r+1 {
				path = append(path, np)
			}
			if p.r == np.r-1 {
				path = append(path, np)
			}
		case "-":
			if p.c == np.c+1 {
				path = append(path, np)
			}
			if p.c == np.c-1 {
				path = append(path, np)
			}
		case "L":
			if np.c == p.c-1 || np.r == p.r+1 {
				path = append(path, np)
			}
		case "J":
			if np.c == p.c+1 || np.r == p.r+1 {
				path = append(path, np)
			}
		case "F":
			if np.c == p.c-1 || np.r == p.r-1 {
				path = append(path, np)
			}
		case "7":
			if np.c == p.c+1 || np.r == p.r-1 {
				path = append(path, np)
			}
		}
		if len(path) > 1 {
			break
		}
	}
	if len(path) != 2 {
		panic("start not found")
	}
	return path
}

// Return new path and then looplen when back to start
// The latest element of path is a neighbor of the previous element
// We know that it was reached from the previous element
func traverse(g u.CharGrid, path []pos) int {
	p := path[len(path)-1]
	printPath(g, path)
	nps := []pos{pos{p.r - 1, p.c}, pos{p.r + 1, p.c}, pos{p.r, p.c - 1}, pos{p.r, p.c + 1}}
	prev := path[len(path)-2]
	c := g.Grid[p.r][p.c]
	longestLoop := 0
	for _, np := range nps {
		if !g.InBounds(np.r, np.c) || np == prev {
			continue
		}
		nc := g.Grid[np.r][np.c]
		if nc == "." {
			continue
		}
		if nc == "S" {
			fmt.Printf("found path to end %d\n", len(path))
			path = append(path, np)
			return len(path)
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
			length := traverse(g, path)
			if length > longestLoop {
				longestLoop = length
			}
		}
	}
	return longestLoop
}

func printPath(g u.CharGrid, path []pos) {
	str := ""
	for _, p := range path {
		str += g.Grid[p.r][p.c]
	}
	fmt.Println(len(str))
}
