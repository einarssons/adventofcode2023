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
	d := u.CreateZeroDigitGrid(g.Width, g.Height)
	pos := pair{0, -1}
	dir := pair{0, 1}
	propagateLight(g, d, pos, dir)
	fmt.Println(d.String())
	return calcEnergy(d)
}

func task2(lines []string) int {
	g := u.CreateRuneGridFromLines(lines)
	maxEnergy := 0

	dir := pair{0, 1}
	for r := 0; r < g.Height; r++ {
		startPos := pair{r, -1}
		d := u.CreateZeroDigitGrid(g.Width, g.Height)
		propagateLight(g, d, startPos, dir)
		energy := calcEnergy(d)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	dir = pair{0, -1}
	for r := 0; r < g.Height; r++ {
		startPos := pair{r, g.Width}
		d := u.CreateZeroDigitGrid(g.Width, g.Height)
		propagateLight(g, d, startPos, dir)
		energy := calcEnergy(d)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	dir = pair{1, 0}
	for c := 0; c < g.Width; c++ {
		startPos := pair{-1, c}
		d := u.CreateZeroDigitGrid(g.Width, g.Height)
		propagateLight(g, d, startPos, dir)
		energy := calcEnergy(d)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	dir = pair{-1, 0}
	for c := 0; c < g.Width; c++ {
		startPos := pair{g.Height, c}
		d := u.CreateZeroDigitGrid(g.Width, g.Height)
		propagateLight(g, d, startPos, dir)
		energy := calcEnergy(d)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	return maxEnergy
}

func calcEnergy(d u.DigitGrid) int {
	count := 0
	for r := 0; r < d.Height; r++ {
		for c := 0; c < d.Width; c++ {
			if d.Grid[r][c] != 0 {
				count++
			}
		}
	}
	return count
}

type pair struct {
	r, c int
}

const (
	up    = 1
	down  = 2
	left  = 4
	right = 8
)

var (
	upP    = pair{-1, 0}
	downP  = pair{1, 0}
	leftP  = pair{0, -1}
	rightP = pair{0, 1}
)

func (p pair) add(p2 pair) pair {
	return pair{p.r + p2.r, p.c + p2.c}
}

// addDir adds a direction to the digit grid and returns true if the value changed
func addDir(pos pair, dir pair, d u.DigitGrid) bool {
	bit := 0
	switch dir {
	case pair{1, 0}:
		bit = down
	case pair{-1, 0}:
		bit = up
	case pair{0, 1}:
		bit = right
	case pair{0, -1}:
		bit = left
	default:
		panic("Invalid direction")
	}
	return addBitDir(pos, bit, d)
}

// addBitDir adds a bit to the digit grid and returns true if the value changed
func addBitDir(pos pair, bit int, d u.DigitGrid) (change bool) {
	old := d.Grid[pos.r][pos.c]
	d.Grid[pos.r][pos.c] |= bit
	return old != d.Grid[pos.r][pos.c]
}

// bitDir returns the bit corresponding to the direction
func bitDir(dir pair) int {
	switch dir {
	case downP:
		return down
	case upP:
		return up
	case rightP:
		return right
	case leftP:
		return left
	default:
		panic("Invalid direction")
	}
}

func propagateLight(g u.RuneGrid, d u.DigitGrid, pos, dir pair) {
	for {
		newPos := pos.add(dir)
		if !g.InBounds(newPos.r, newPos.c) {
			return
		}
		bit := bitDir(dir)
		switch g.Grid[newPos.r][newPos.c] {
		case '.':
			addDir(newPos, dir, d)
		case '|':
			if bit == left || bit == right {
				// Split up/down
				if addBitDir(newPos, bit|up, d) {
					propagateLight(g, d, newPos, upP)
				}
				if addBitDir(newPos, bit|down, d) {
					propagateLight(g, d, newPos, downP)
				}
				return
			} else {
				addDir(newPos, dir, d)
			}
		case '-':
			if bit == up || bit == down {
				// Split left/right
				if addBitDir(newPos, bit|left, d) {
					propagateLight(g, d, newPos, leftP)
				}
				if addBitDir(newPos, bit|right, d) {
					propagateLight(g, d, newPos, rightP)
				}
				return
			} else {
				addDir(newPos, dir, d)
			}
		case '\\':
			switch bit {
			case right:
				addBitDir(newPos, bit|down, d)
				dir = downP
			case up:
				addBitDir(newPos, bit|left, d)
				dir = leftP
			case left:
				addBitDir(newPos, bit|up, d)
				dir = upP
			case down:
				addBitDir(newPos, bit|right, d)
				dir = rightP
			default:
				panic("Invalid direction")
			}
		case '/':
			switch bit {
			case right:
				addBitDir(newPos, bit|up, d)
				dir = upP
			case up:
				addBitDir(newPos, bit|right, d)
				dir = rightP
			case left:
				addBitDir(newPos, bit|down, d)
				dir = downP
			case down:
				addBitDir(newPos, bit|left, d)
				dir = leftP
			default:
				panic("Invalid direction")
			}
		}
		pos = newPos
	}
}
