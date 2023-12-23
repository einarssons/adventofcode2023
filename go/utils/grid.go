package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Grid2D[T comparable] struct {
	Grid   [][]T
	Width  int
	Height int
}

func CreateGrid2D[T comparable](width, height int) Grid2D[T] {
	grid := Grid2D[T]{
		Grid:   make([][]T, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]T, grid.Width))
	}
	return grid
}

// InBounds checks if (row, col) inside grid
func (g Grid2D[T]) InBounds(row, col int) bool {
	return 0 <= row && row < g.Height && 0 <= col && col < g.Width
}

// AtBorder checks if (row, col) is at the border of the grid
func (g Grid2D[T]) AtBorder(row, col int) bool {
	return row == 0 || row == g.Height-1 || col == 0 || col == g.Width-1
}

// SetAll sets all cells in the grid to the given value
func (g *Grid2D[T]) SetAll(value T) {
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			g.Grid[row][col] = value
		}
	}
}

func (g *Grid2D[T]) Find(value T) (row, col int, ok bool) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == value {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}

func (g *Grid2D[T]) At(row, col int) T {
	return g.Grid[row][col]
}

// Set sets an individual cell in the grid
func (g *Grid2D[T]) Set(val T, row, col int) {
	g.Grid[row][col] = val
}

func (g Grid2D[T]) Get(row, col int) T {
	return g.Grid[row][col]
}

type DigitGrid struct {
	Grid   [][]int
	Width  int
	Height int
}

func CreateDigitGridFromLines(lines []string) DigitGrid {
	g := DigitGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := make([]int, 0, g.Width)
		digits := SplitToChars(line)
		for _, digit := range digits {
			nr, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			row = append(row, nr)
		}
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

func (g *DigitGrid) String() string {
	var rows []string
	for r := 0; r < g.Height; r++ {
		row := ""
		for c := 0; c < g.Width; c++ {
			row += fmt.Sprintf("%x", g.Grid[r][c])
		}
		rows = append(rows, row)
	}
	return "\n" + strings.Join(rows, "\n") + "\n"
}

type CharGrid struct {
	Grid   [][]string
	Width  int
	Height int
}

func CreateCharGridFromLines(lines []string) CharGrid {
	g := CharGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := SplitToChars(line)
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

func CreateEmptyCharGrid(width, height int) CharGrid {
	grid := CharGrid{
		Grid:   make([][]string, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]string, grid.Width))
	}
	return grid
}

func (g *CharGrid) SetAll(value string) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Grid[y][x] = value
		}
	}
}

func (g *CharGrid) Find(x string) (row, col int) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == x {
				return r, c
			}
		}
	}
	return -1, -1
}

func (g CharGrid) String() string {
	var rows []string
	for r := 0; r < g.Height; r++ {
		row := strings.Join(g.Grid[r], "")
		rows = append(rows, row)
	}
	return "\n" + strings.Join(rows, "\n") + "\n"
}

// InBounds - is (y, x) in grid
func (g CharGrid) InBounds(y, x int) bool {
	return 0 <= y && y < g.Height && 0 <= x && x < g.Width
}

func CreateZeroDigitGrid(width, height int) DigitGrid {
	grid := DigitGrid{
		Grid:   make([][]int, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]int, grid.Width))
	}
	return grid
}

// InBounds - is (y, x) in grid
func (g DigitGrid) InBounds(y, x int) bool {
	return 0 <= y && y < g.Height && 0 <= x && x < g.Width
}

func (g *DigitGrid) SetAll(value int) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Grid[y][x] = value
		}
	}
}

type RuneGrid struct {
	Grid   [][]rune
	Width  int
	Height int
}

func CreateRuneGridFromLines(lines []string) RuneGrid {
	g := RuneGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := SplitToRunes(line)
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

func (g *RuneGrid) SetAll(value rune) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			g.Grid[r][c] = value
		}
	}
}

func (g *RuneGrid) Find(x rune) (row, col int) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == x {
				return r, c
			}
		}
	}
	return -1, -1
}

func (g RuneGrid) String() string {
	var rows []string
	for r := 0; r < g.Height; r++ {
		row := string(g.Grid[r])
		rows = append(rows, row)
	}
	return "\n" + strings.Join(rows, "\n") + "\n"
}

// InBounds - is (row, col) in grid
func (g RuneGrid) InBounds(r, c int) bool {
	return 0 <= r && r < g.Height && 0 <= c && c < g.Width
}

func (g RuneGrid) At(r, c int) rune {
	return g.Grid[r][c]
}

func (g RuneGrid) Copy() RuneGrid {
	n := RuneGrid{
		Grid:   make([][]rune, 0, g.Height),
		Width:  g.Width,
		Height: g.Height,
	}
	for r := 0; r < g.Height; r++ {
		row := make([]rune, g.Width)
		copy(row, g.Grid[r])
		n.Grid = append(n.Grid, row)
	}
	return n
}
