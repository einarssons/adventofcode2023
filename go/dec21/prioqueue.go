package main

import (
	"container/heap"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

type Item struct {
	pos   u.Pos2D
	steps int
	index int // Internal index in the heap
}

func (i Item) priority() int {
	return i.steps
}

func NewItem(pos u.Pos2D, steps int) *Item {
	return &Item{pos: pos, steps: steps}
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority() < pq[j].priority()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, pos u.Pos2D, steps int) {
	item.pos = pos
	item.steps = steps
	heap.Fix(pq, item.index)
}
