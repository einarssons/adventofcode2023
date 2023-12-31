package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

func Test(t *testing.T) {
	g := Graph{
		Vertices: []Vertice{
			{ID: 1, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 2, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 3, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 4, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 5, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 6, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 7, OriginalIDs: u.CreateSet[VerticeName]()},
			{ID: 8, OriginalIDs: u.CreateSet[VerticeName]()},
		},
		Edges: []Edge{
			{From: 1, To: 2, Cost: 2},
			{From: 2, To: 3, Cost: 3},
			{From: 3, To: 4, Cost: 4},
			{From: 1, To: 5, Cost: 3},
			{From: 2, To: 5, Cost: 2},
			{From: 2, To: 6, Cost: 2},
			{From: 3, To: 7, Cost: 2},
			{From: 4, To: 7, Cost: 2},
			{From: 4, To: 8, Cost: 2},
			{From: 5, To: 6, Cost: 3},
			{From: 6, To: 7, Cost: 1},
			{From: 7, To: 8, Cost: 3},
		},
	}
	for i := range g.Vertices {
		g.Vertices[i].OriginalIDs.Add(VerticeName(fmt.Sprintf("%d", i+1)))
	}
	cut := MinimumCut(g)
	fmt.Println("Minimal cut: ", cut)
	if cut.Cost != 4 {
		t.Errorf("Expected cost 4, got %d", cut.Cost)
	}
}
