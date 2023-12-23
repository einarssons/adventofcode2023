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
	steps := parse(lines)
	h := holes{cut: make(map[Pos2D]string)}
	pos := Pos2D{0, 0}
	h.cut[pos] = ""
	for _, s := range steps {
		var dir Pos2D
		switch s.dir {
		case "R":
			dir = Pos2D{0, 1}
		case "L":
			dir = Pos2D{0, -1}
		case "U":
			dir = Pos2D{-1, 0}
		case "D":
			dir = Pos2D{1, 0}
		}
		for i := 0; i < s.steps; i++ {
			pos = pos.add(dir)
			if _, ok := h.cut[pos]; ok {
				fmt.Println("Collision at", pos)
			}
			h.cut[pos] = s.color
		}
	}
	minRow, maxRow, minCol, maxCol := h.findBounds()
	nrExterior := h.markExterior(minRow, maxRow, minCol, maxCol)
	volume := (maxRow-minRow+1)*(maxCol-minCol+1) - nrExterior
	return volume
}

func task2(lines []string) int {
	steps := parse2(lines)
	pos := Pos2D{0, 0}
	startPos := pos
	corners := make([]Pos2D, 0, len(steps))
	corners = append(corners, pos)
	closed := false
	edgeLen := 0
	for i, s := range steps {
		var dir Pos2D
		switch s.dir {
		case "R":
			dir = Pos2D{0, 1}
		case "L":
			dir = Pos2D{0, -1}
		case "U":
			dir = Pos2D{-1, 0}
		case "D":
			dir = Pos2D{1, 0}
		}
		edgeLen += s.steps
		pos = pos.add(dir.mul(s.steps))
		if closed {
			panic(fmt.Sprintf("closed but step=%d (%d)", i, len(steps)))
		}
		if pos == startPos {
			closed = true
			fmt.Println("back to start")
			continue
		}
		corners = append(corners, pos)
	}
	area := shoelace(corners) + edgeLen/2 + 1
	return area
}

func shoelace(corners []Pos2D) int {
	area := 0
	for i := 0; i < len(corners)-1; i++ {
		iplus := (i + 1) % len(corners)
		area += (corners[i].row + corners[iplus].row) * (corners[i].col - corners[iplus].col)
	}
	return u.Abs(area) / 2
}

type step struct {
	dir   string
	steps int
	color string
}

type Pos2D struct {
	row, col int
}

func (p Pos2D) neg() Pos2D {
	return Pos2D{-p.row, -p.col}
}

func (p Pos2D) sub(other Pos2D) Pos2D {
	return Pos2D{p.row - other.row, p.col - other.col}
}

func (p Pos2D) add(other Pos2D) Pos2D {
	return Pos2D{p.row + other.row, p.col + other.col}
}

func (p Pos2D) mul(factor int) Pos2D {
	return Pos2D{factor * p.row, factor * p.col}
}

// left is a 90 degree left rotation of the vector
func (p Pos2D) left() Pos2D {
	return Pos2D{-p.col, p.row}
}

// right is a 90 degree right rotation of the vector
func (p Pos2D) right() Pos2D {
	return Pos2D{p.col, -p.row}
}

type holes struct {
	cut map[Pos2D]string
}

type bounds struct {
	minRow, maxRow, minCol, maxCol int
}

func (h *holes) markExterior(minRow, maxRow, minCol, maxCol int) int {
	nrMarked := 0
	for row := minRow; row <= maxRow; row++ {
		pos := Pos2D{row, minCol}
		if _, ok := h.cut[pos]; !ok {
			nrMarked += h.markNeighbors(pos, minRow, maxRow, minCol, maxCol)
		}
		pos = Pos2D{row, maxCol}
		if _, ok := h.cut[pos]; !ok {
			nrMarked += h.markNeighbors(pos, minRow, maxRow, minCol, maxCol)
		}
	}
	for col := minCol; col <= maxCol; col++ {
		pos := Pos2D{minRow, col}
		if _, ok := h.cut[pos]; !ok {
			nrMarked += h.markNeighbors(pos, minRow, maxRow, minCol, maxCol)
		}
		pos = Pos2D{maxRow, col}
		if _, ok := h.cut[pos]; !ok {
			nrMarked += h.markNeighbors(pos, minRow, maxRow, minCol, maxCol)
		}
	}
	return nrMarked
}

func (h *holes) markNeighbors(pos Pos2D, minRow, maxRow, minCol, maxCol int) int {
	nrMarked := 1
	h.cut[pos] = "#"
	for _, dir := range []Pos2D{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		n := pos.add(dir)
		if n.row < minRow || n.row > maxRow || n.col < minCol || n.col > maxCol {
			continue
		}
		if _, ok := h.cut[n]; !ok {
			nrMarked += h.markNeighbors(n, minRow, maxRow, minCol, maxCol)
		}
	}
	return nrMarked
}

func (h *holes) findBounds() (minRow, maxRow, minCol, maxCol int) {
	minRow, maxRow, minCol, maxCol = u.MaxInt, -u.MaxInt, u.MaxInt, -u.MaxInt
	for pos := range h.cut {
		if pos.row < minRow {
			minRow = pos.row
		}
		if pos.row > maxRow {
			maxRow = pos.row
		}
		if pos.col < minCol {
			minCol = pos.col
		}
		if pos.col > maxCol {
			maxCol = pos.col
		}

	}
	return minRow, maxRow, minCol, maxCol
}

func (h *holes) countWallsLeft(minCol, row, col int) int {
	nrLeft := 0
	inWall := false
	for c := minCol; c < col; c++ {
		color, ok := h.cut[Pos2D{row, c}]
		if !ok || color == "#" {
			if inWall {
				nrLeft++
				inWall = false
			}
		}
		if ok && color != "#" {
			inWall = true
		}
	}
	if inWall {
		nrLeft++
	}
	return nrLeft
}

func (h *holes) countWallsRight(maxCol, row, col int) int {
	nrRight := 0
	inWall := false
	for c := maxCol; c > col; c-- {
		color, ok := h.cut[Pos2D{row, c}]
		if !ok || color == "#" {
			if inWall {
				nrRight++
				inWall = false
			}
		}
		if ok && color != "#" {
			inWall = true
		}
	}
	if inWall {
		nrRight++
	}
	return nrRight
}

func (h *holes) countWallsTop(minRow, row, col int) int {
	nrTop := 0
	inWall := false
	for r := minRow; r < row; r++ {
		color, ok := h.cut[Pos2D{r, col}]
		if !ok || color == "#" {
			if inWall {
				nrTop++
				inWall = false
			}
		}
		if ok && color != "#" {
			inWall = true
		}
	}
	if inWall {
		nrTop++
	}
	return nrTop
}

func (h *holes) countWallsBottom(maxRow, row, col int) int {
	nrBottom := 0
	inWall := false
	for r := maxRow; r > row; r-- {
		color, ok := h.cut[Pos2D{r, col}]
		if !ok || color == "#" {
			if inWall {
				nrBottom++
				inWall = false
			}
		}
		if ok && color != "#" {
			inWall = true
		}
	}
	if inWall {
		nrBottom++
	}
	return nrBottom
}

func parse(lines []string) []step {
	steps := make([]step, 0, len(lines))
	for _, line := range lines {
		dir := line[0:1]
		p, c := u.Cut(line[2:], " ")
		s := u.Atoi(p)
		col := c[1 : len(c)-1]
		steps = append(steps, step{dir, s, col})
	}
	return steps
}

func parse2(lines []string) []step {
	steps := make([]step, 0, len(lines))
	for _, line := range lines {
		ll := len(line)
		d := line[ll-2 : ll-1]
		var dir string
		switch d {
		case "0":
			dir = "R"
		case "1":
			dir = "D"
		case "2":
			dir = "L"
		case "3":
			dir = "U"
		default:
			panic("Unknown direction")
		}
		hexDigits := line[ll-7 : ll-2]
		s, err := strconv.ParseInt(hexDigits, 16, 64)
		if err != nil {
			panic(err)
		}
		steps = append(steps, step{dir, int(s), ""})
	}
	return steps
}
