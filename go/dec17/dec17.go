package main

import (
	"container/heap"
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

type state struct {
	pos   Pos2D
	dir   Pos2D
	steps int
}

func task1(lines []string) int {
	g := u.CreateDigitGridFromLines(lines)
	endPos := Pos2D{g.Height - 1, g.Width - 1}
	result := traverse(g, endPos)
	return result
}

func traverse(g u.DigitGrid, endPos Pos2D) int {
	endDist := func(pos Pos2D) int {
		diff := pos.sub(endPos)
		return u.Abs(diff.row) + u.Abs(diff.col)
	}
	newItem := func(pos, dir Pos2D, cost, steps int) *Item {
		return &Item{pos: pos, dir: dir, cost: cost, steps: steps, dist: endDist(pos)}
	}
	pq := make(PriorityQueue, 0, 100)
	pq.Push(newItem(Pos2D{0, 0}, Pos2D{1, 0}, 0, 0))
	pq.Push(newItem(Pos2D{0, 0}, Pos2D{0, 1}, 0, 0))
	heap.Init(&pq)
	visited := u.CreateSet[state]()
	for {
		if pq.Len() == 0 {
			break
		}
		curr := heap.Pop(&pq).(*Item)
		currState := state{curr.pos, curr.dir, curr.steps}
		if visited.Contains(currState) {
			continue
		}
		visited.Add(currState)
		currPos := Pos2D{curr.pos.row, curr.pos.col}
		if currPos == endPos {
			fmt.Println("found end at cost ", curr.cost)
			return curr.cost
		}
		if curr.steps < 3 {
			newPos := currPos.add(curr.dir)
			if g.InBounds(newPos.row, newPos.col) {
				pq.Push(newItem(newPos, curr.dir,
					curr.cost+g.Grid[newPos.row][newPos.col],
					curr.steps+1))
			}
		}
		newDir := curr.dir.left()
		newPos := currPos.add(newDir)
		if g.InBounds(newPos.row, newPos.col) {
			pq.Push(newItem(newPos, newDir,
				curr.cost+g.Grid[newPos.row][newPos.col], 1))
		}
		newDir = curr.dir.right()
		newPos = currPos.add(newDir)
		if g.InBounds(newPos.row, newPos.col) {
			pq.Push(newItem(newPos, newDir,
				curr.cost+g.Grid[newPos.row][newPos.col], 1))
		}
	}
	panic("no path found")
}

func task2(lines []string) int {
	g := u.CreateDigitGridFromLines(lines)
	endPos := Pos2D{g.Height - 1, g.Width - 1}
	result := traverse2(g, endPos)
	return result
}

func traverse2(g u.DigitGrid, endPos Pos2D) int {
	endDist := func(pos Pos2D) int {
		diff := pos.sub(endPos)
		return u.Abs(diff.row) + u.Abs(diff.col)
	}
	newItem := func(pos, dir Pos2D, cost, steps int) *Item {
		return &Item{pos: pos, dir: dir, cost: cost, steps: steps, dist: endDist(pos)}
	}
	pq := make(PriorityQueue, 0, 100)
	pq.Push(newItem(Pos2D{0, 0}, Pos2D{1, 0}, 0, 0))
	pq.Push(newItem(Pos2D{0, 0}, Pos2D{0, 1}, 0, 0))
	heap.Init(&pq)
	visited := u.CreateSet[state]()
	for {
		if pq.Len() == 0 {
			break
		}
		curr := heap.Pop(&pq).(*Item)
		currState := state{curr.pos, curr.dir, curr.steps}
		if visited.Contains(currState) {
			continue
		}
		visited.Add(currState)
		currPos := Pos2D{curr.pos.row, curr.pos.col}
		if currPos == endPos && curr.steps > 3 {
			fmt.Println("found end at cost ", curr.cost)
			return curr.cost
		}
		if curr.steps < 10 { // max straight steps
			newPos := currPos.add(curr.dir)
			if g.InBounds(newPos.row, newPos.col) {
				pq.Push(newItem(newPos, curr.dir,
					curr.cost+g.Grid[newPos.row][newPos.col],
					curr.steps+1))
			}
			if curr.steps < 4 {
				continue // Can't turn yet
			}
		}
		newDir := curr.dir.left()
		newPos := currPos.add(newDir)
		if g.InBounds(newPos.row, newPos.col) {
			pq.Push(newItem(newPos, newDir,
				curr.cost+g.Grid[newPos.row][newPos.col], 1))
		}
		newDir = curr.dir.right()
		newPos = currPos.add(newDir)
		if g.InBounds(newPos.row, newPos.col) {
			pq.Push(newItem(newPos, newDir,
				curr.cost+g.Grid[newPos.row][newPos.col], 1))
		}
	}
	panic("no path found")
}
