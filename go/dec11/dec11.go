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
		fmt.Println("task1: ", task(lines, 2))
	} else {
		fmt.Println("task2: ", task(lines, 1_000_000))
	}
}

type coord struct {
	r, c int
}

func task(lines []string, factor int) int {
	grid := u.CreateCharGridFromLines(lines)
	galaxies := findGalaxies(grid)
	emtyRows := findEmptyRows(grid)
	emtyCols := findEmptyCols(grid)

	sum := 0
	nrGalaxies := len(galaxies)
	for i := 0; i < nrGalaxies; i++ {
		for j := i + 1; j < nrGalaxies; j++ {
			dist := u.Abs(galaxies[i].r-galaxies[j].r) + u.Abs(galaxies[i].c-galaxies[j].c)
			dist += nrRowsBetween(galaxies[i], galaxies[j], emtyRows) * (factor - 1)
			dist += nrColsBetween(galaxies[i], galaxies[j], emtyCols) * (factor - 1)
			sum += dist
		}
	}
	return sum
}

func findGalaxies(grid u.CharGrid) []coord {
	galaxies := []coord{}
	for r := 0; r < grid.Height; r++ {
		for c := 0; c < grid.Width; c++ {
			if grid.Grid[r][c] == "#" {
				galaxies = append(galaxies, coord{r, c})
			}
		}
	}
	return galaxies
}

func findEmptyRows(grid u.CharGrid) u.Set[int] {
	emptyRows := u.CreateSet[int]()
	for r := 0; r < grid.Height; r++ {
		empty := true
		for c := 0; c < grid.Width; c++ {
			if grid.Grid[r][c] == "#" {
				empty = false
				break
			}
		}
		if empty {
			emptyRows.Add(r)
		}
	}
	return emptyRows
}

func findEmptyCols(grid u.CharGrid) u.Set[int] {
	emptyCols := u.CreateSet[int]()
	for c := 0; c < grid.Width; c++ {
		empty := true
		for r := 0; r < grid.Height; r++ {
			if grid.Grid[r][c] == "#" {
				empty = false
				break
			}
		}
		if empty {
			emptyCols.Add(c)
		}
	}
	return emptyCols
}

func nrRowsBetween(c1, c2 coord, emptyRows u.Set[int]) int {
	if c1.r > c2.r {
		c1, c2 = c2, c1
	}
	nr := 0
	for r := c1.r + 1; r < c2.r; r++ {
		if emptyRows.Contains(r) {
			nr++
		}
	}
	return nr
}

func nrColsBetween(c1, c2 coord, emptyCols u.Set[int]) int {
	if c1.c > c2.c {
		c1, c2 = c2, c1
	}
	nr := 0
	for c := c1.c + 1; c < c2.c; c++ {
		if emptyCols.Contains(c) {
			nr++
		}
	}
	return nr
}
