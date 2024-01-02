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
		fmt.Println("task1: ", task1(lines, 64))
	} else {
		fmt.Println("task2: ", task2(lines, 26501365))
	}
}

func task1(lines []string, maxSteps int) int {
	g := parse(lines)
	fmt.Printf("maxSteps %d, grid is %dx%d and start at %d,%d\n", maxSteps, g.width, g.height, g.start.Row, g.start.Col)
	if g.width != g.height {
		panic("grid is not square")
	}
	plots := make(map[u.Pos2D]plot)
	pq := make(PriorityQueue, 0, 100)
	walk(g, plots, pq, maxSteps)
	nrPlots := 0
	for _, p := range plots {
		if p.steps%2 == 0 {
			nrPlots++
		}
	}
	return nrPlots
}

type plot struct {
	steps int
	done  bool
}

func (p plot) oddSteps() bool {
	return p.steps%2 == 1
}

func task2(lines []string, maxSteps int) int {
	g := parse(lines)
	fmt.Printf("maxSteps %d, grid is %dx%d and start at %d,%d\n", maxSteps, g.width, g.height, g.start.Row, g.start.Col)
	if g.width != g.height {
		panic("grid is not square")
	}
	side := g.width
	maxTileNr := (maxSteps - 65) / side
	plots := make(map[u.Pos2D]plot)
	pq := make(PriorityQueue, 0, 100)
	// Any even value of nrTilesToWalk generate 13 different
	// numbers of odd interior plots. We need to find the right
	// combination to scale to maxSteps.
	nrTilesToWalk := 2
	maxStepsToWalk := 65 + nrTilesToWalk*side
	fmt.Println("maxStepsToWalk: ", maxStepsToWalk)
	walk(g, plots, pq, maxStepsToWalk)
	// Next use the diamond shape to find a repeatable number of plots
	gridCounts := make(map[u.Pos2D]int)
	for gridRow := -nrTilesToWalk; gridRow <= nrTilesToWalk; gridRow++ {
		for gridCol := -nrTilesToWalk; gridCol <= nrTilesToWalk; gridCol++ {
			count := 0
			for row := gridRow * side; row < (gridRow+1)*side; row++ {
				for col := gridCol * side; col < (gridCol+1)*side; col++ {
					if steps, ok := plots[u.Pos2D{row, col}]; ok {
						if steps.oddSteps() && steps.done {
							count++
						}
					}
				}
			}
			gridCounts[u.Pos2D{Row: gridRow, Col: gridCol}] = count
		}
	}
	counts := make(map[int][]u.Pos2D)
	for gridRow := -nrTilesToWalk; gridRow <= nrTilesToWalk; gridRow++ {
		for gridCol := -nrTilesToWalk; gridCol <= nrTilesToWalk; gridCol++ {
			count := gridCounts[u.Pos2D{Row: gridRow, Col: gridCol}]
			if count != 0 {
				p := u.Pos2D{Row: gridRow, Col: gridCol}
				counts[count] = append(counts[count], p)
			}
		}
	}
	i := 0
	for k, v := range counts {
		fmt.Printf("%d  %4d %v\n", i, k, v)
		i++
	}
	/* Next need to add upp all number of odd interior tiles, even interior tiles
	   edge and corner tiles. For 5x5 grid (maxTileNr = 2) we get:

		   0   995 [{-2 1} {-1 2}] ERTE x n
		   1  5799 [{0 2}] ORGT x 1
		   2   991 [{1 2} {2 1}] ERBE x n
		   3  6778 [{-1 -1}] OLTE x (n-1)
		   4  5800 [{0 -2}] OLFT x 1
		   5  6744 [{1 -1}] Olbe x (n-1)
		   6  6753 [{1 1}] ORBE x (n-1)
		   7  1012 [{-2 -1} {-1 -2}] ELTE x n
		   8  7722 [{0 0}] EVEN x (n-1)*(n-1)
		   9   988 [{1 -2} {2 -1}] ELBE x n
		   10  5824 [{-2 0}] ODDT x 1
		   11  7753 [{-1 0} {0 -1} {0 1} {1 0}] EVEN x n*n
		   12  6768 [{-1 1}] ORTE x (n-1)
		   13  5775 [{2 0}] ODDB x 1

		      -2   -1   0    1    2
		-2   ---- ELTE ODDT ERTE ----
		-1   ELTE OLTE EVEN ORTE ERTE
		 0   OLFT EVEN ODDF EVEN ORGT
		+1   ELBE OLBE EVEN ORBE ERBE
		+2   ---- ELBE ODDB ERBE ----

	*/

	n := maxTileNr
	nrOddFull := (n - 1) * (n - 1) // [0,0]
	nrEvenFull := n * n            // [0,1]
	nrOddEdges := n - 1            // [-1, -1], [-1, 1], [1, -1], [1, 1]
	nrEvenEdges := n               // [1, -2], [1, 2], [-1, -2], [-1, 2]
	// nrCorners = 1 (0,2), 1 (2,0), 1 (0,-2), 1 (-2,0)

	c := func(r, c int) int {
		return gridCounts[u.Pos2D{Row: r, Col: c}]
	}

	nrPlots := nrOddFull*c(0, 0) + nrEvenFull*c(0, 1) +
		nrOddEdges*(c(-1, 1)+c(-1, -1)+c(1, -1)+c(1, 1)) +
		nrEvenEdges*(c(-1, -2)+c(-1, 2)+c(1, -2)+c(1, 2)) +
		c(2, 0) + c(0, 2) + c(-2, 0) + c(0, -2)

	return nrPlots
}

func walk(g gardenTile, plots map[u.Pos2D]plot, pq PriorityQueue, maxSteps int) {
	dirs := []u.Pos2D{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	start := Item{pos: g.start, steps: 0}
	pq.Push(&start)
	for {
		if pq.Len() == 0 {
			break
		}
		item := heap.Pop(&pq).(*Item)
		plots[item.pos] = plot{steps: item.steps, done: true}
		newSteps := item.steps + 1
		if newSteps > maxSteps {
			continue
		}
		for _, d := range dirs {
			n := item.pos.Add(d)
			if g.inPath(n) {
				if _, ok := plots[n]; !ok {
					plots[n] = plot{steps: newSteps, done: false}
					pq.Push(NewItem(n, newSteps))
				}
			}
		}
	}
}

// gardenTile is the repeated map tile in the garden
type gardenTile struct {
	s      u.Set[u.Pos2D]
	start  u.Pos2D
	width  int
	height int
}

func (g gardenTile) inPath(pos u.Pos2D) bool {
	r := pos.Row % g.height
	c := pos.Col % g.width
	if r < 0 {
		r += g.height
	}
	if c < 0 {
		c += g.width
	}
	return g.s.Contains(u.Pos2D{r, c})
}

func parse(lines []string) gardenTile {
	t := gardenTile{
		width:  len(lines[0]),
		height: len(lines),
		s:      u.CreateSet[u.Pos2D](),
	}
	nrOdd := 0
	for row, line := range lines {
		for col, c := range line {
			switch c {
			case '#':
				continue
			case '.':
				t.s.Add(u.Pos2D{row, col})
				if (row+col)%2 == 1 {
					nrOdd++
				}
			case 'S':
				p := u.Pos2D{row, col}
				t.s.Add(p)
				t.start = p
				if (row+col)%2 == 1 {
					nrOdd++
				}
			default:
				panic("unknown char")
			}
		}
	}
	return t
}
