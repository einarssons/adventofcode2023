package main

import (
	"flag"
	"fmt"
	"strconv"

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
	grid := u.CreateCharGridFromLines(lines)
	numbers := findNumbers(grid, true)
	sum := 0
	for _, n := range numbers {
		sum += n.val
	}
	return sum
}

func task2(lines []string) int {
	grid := u.CreateCharGridFromLines(lines)
	numbers := findGearRatios(grid)
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

type gridNr struct {
	row, col, val int
}

// endCol returns the column index of the last digit of the number
func (g gridNr) endCol() int {
	return g.col + len(strconv.Itoa(g.val)) - 1
}

func findNumbers(grid u.CharGrid, withNeighborSymbol bool) []gridNr {
	var nrs []gridNr
	for row, line := range grid.Grid {
		w := len(line)
		col := 0
		for {
			if isDigit(line[col]) {
				gNr := gridNr{row, col, 0}
				start := col
				end := col + 1
				nr := u.Atoi(line[col])
				col++
				for {
					if col >= w {
						break
					}
					if isDigit(line[col]) {
						nr = nr*10 + u.Atoi(line[col])
						end = col + 1
					} else {
						break
					}
					col++
				}
				gNr.val = nr

				if withNeighborSymbol {
					if hasNeighborSymbol(grid, row, start, end) {
						nrs = append(nrs, gNr)
						//fmt.Println("Found number", nr, "at", row, start, end)
					}
				} else {
					nrs = append(nrs, gNr)
				}
				col = end
			} else {
				col++
			}
			if col >= w {
				break
			}
		}
	}
	return nrs
}

func findGearRatios(grid u.CharGrid) []int {
	ratios := make([]int, 0)
	gridNrs := findNumbers(grid, false)
	for row, line := range grid.Grid {
		for col := range line {
			if grid.Grid[row][col] == "*" {
				neighbors := findAdjacentNumbers(grid, row, col, gridNrs)
				if len(neighbors) == 2 {
					ratios = append(ratios, neighbors[0]*neighbors[1])
				}
			}
		}
	}
	return ratios
}

func findAdjacentNumbers(grid u.CharGrid, row, col int, gridNrs []gridNr) []int {
	nrs := u.CreateSet[int]()
	for r := row - 1; r <= row+1; r += 2 {
		for _, gNr := range gridNrs {
			if gNr.row == r {
				if gNr.endCol() >= col-1 && gNr.col <= col+1 {
					nrs.Add(gNr.val)
				}
			}
		}
	}
	for _, gNr := range gridNrs { // Same line
		if gNr.row == row {
			if gNr.endCol() == col-1 || gNr.col == col+1 {
				nrs.Add(gNr.val)
			}
		}
	}
	return nrs.Values()
}

func isDigit(s string) bool {
	return s >= "0" && s <= "9"
}

func isSymbol(grid u.CharGrid, r, c int) bool {
	return grid.InBounds(r, c) && grid.Grid[r][c] != "." && !isDigit(grid.Grid[r][c])
}

func hasNeighborSymbol(grid u.CharGrid, row, start, end int) bool {
	for c := start - 1; c < end+1; c++ {
		if isSymbol(grid, row-1, c) {
			return true
		}
	}
	if isSymbol(grid, row, start-1) {
		return true
	}
	if isSymbol(grid, row, end) {
		return true
	}
	for c := start - 1; c < end+1; c++ {
		if isSymbol(grid, row+1, c) {
			return true
		}
	}
	return false
}
