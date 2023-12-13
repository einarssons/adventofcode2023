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
	grids := parse(lines)
	sum := 0
	for i, grid := range grids {
		vert, horiz := reflectionsWithSkip(grid, -1, -1)
		fmt.Printf("%d col=%d row=%d %dx%d\n", i, vert, horiz, grid.Width, grid.Height)
		sum += vert + 100*horiz
	}
	return sum
}

func task2(lines []string) int {
	grids := parse(lines)
	sum := 0
	for i, grid := range grids {
		vert, horiz := reflectionsWithSkip(grid, -1, -1)
		v2, h2 := smudge(grid, vert, horiz)
		fmt.Printf("%d col=%d row=%d %dx%d\n", i, v2, h2, grid.Width, grid.Height)
		sum += v2 + 100*h2
	}
	return sum
}

func flip(grid u.CharGrid, r, c int) {
	switch grid.Grid[r][c] {
	case "#":
		grid.Grid[r][c] = "."
	case ".":
		grid.Grid[r][c] = "#"
	}
}

func smudge(grid u.CharGrid, vert int, horiz int) (v int, h int) {
	for i := 0; i < grid.Height; i++ {
		for j := 0; j < grid.Width; j++ {
			flip(grid, i, j)
			v, h := reflectionsWithSkip(grid, vert, horiz)
			flip(grid, i, j)
			if v != 0 || h != 0 {
				return v, h
			}
		}
	}
	panic("no smudge found")
}

func reflectionsWithSkip(grid u.CharGrid, prevVer, prevHor int) (ver int, hor int) {
	for midRow := 1; midRow <= grid.Height-1; midRow++ {
		symmetric := true
		for d := 1; d < grid.Height; d++ {
			upRow := midRow - d
			downRow := midRow + d - 1
			if upRow < 0 || downRow >= grid.Height {
				break
			}
			if !cmp(grid.Grid[upRow], grid.Grid[downRow]) {
				symmetric = false
				break
			}
		}
		if symmetric && midRow != prevHor {
			hor = midRow
			break
		}
	}
	for midCol := 1; midCol <= grid.Width-1; midCol++ {
		symmetric := true
		for d := 1; d < grid.Width; d++ {
			leftCol := midCol - d
			rightCol := midCol + d - 1
			if leftCol < 0 || rightCol >= grid.Width {
				break
			}
			if !cmpCols(grid, leftCol, rightCol) {
				symmetric = false
				break
			}
		}
		if symmetric && midCol != prevVer {
			ver = midCol
			break
		}
	}

	return ver, hor
}

func cmp(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, s := range s1 {
		if s != s2[i] {
			return false
		}
	}
	return true
}

func cmpCols(grid u.CharGrid, col1 int, col2 int) bool {
	if col1 < 0 || col1 >= grid.Width || col2 < 0 || col2 >= grid.Width {
		return false
	}
	for i := 0; i < grid.Height; i++ {
		if grid.Grid[i][col1] != grid.Grid[i][col2] {
			return false
		}
	}
	return true
}

func parse(lines []string) []u.CharGrid {
	grids := []u.CharGrid{}
	firstRow := 0
	row := 0
	for {
		if row == len(lines) {
			break
		}
		if lines[row] == "" {
			grid := u.CreateCharGridFromLines(lines[firstRow:row])
			grids = append(grids, grid)
			firstRow = row + 1
			row++
			continue
		}
		row++
	}
	grid := u.CreateCharGridFromLines(lines[firstRow:row])
	grids = append(grids, grid)

	return grids
}
