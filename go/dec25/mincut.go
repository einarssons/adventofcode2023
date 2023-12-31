// Stoer-Wagner algorithm for finding the minimum cut in a graph
// https://e-maxx.ru/bookz/files/stoer_wagner_mincut.pdf
package main

import (
	"fmt"
	"math"
	"sort"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

type Graph struct {
	Vertices []Vertice
	Edges    []Edge
}

func (g Graph) String() string {
	s := "Graph:\n"
	for _, v := range g.Vertices {
		s += fmt.Sprintf("  %s\n", v)
	}
	for _, e := range g.Edges {
		s += fmt.Sprintf("  %s\n", e)
	}
	return s
}

func (g Graph) VerticeIDFromName(name VerticeName) VerticeID {
	for _, v := range g.Vertices {
		if v.OriginalIDs.Contains(name) {
			return v.ID
		}
	}
	panic(fmt.Sprintf("Vertice %s not found", name))
}

type VerticeID uint
type VerticeName string

type Edge struct {
	From, To VerticeID
	Cost     uint
}

func (e Edge) String() string {
	return fmt.Sprintf("%d -> %d: %d", e.From, e.To, e.Cost)
}

type Vertice struct {
	ID          VerticeID
	OriginalIDs u.Set[VerticeName]
}

func NewVertice(nr int, name string) Vertice {
	v := Vertice{
		ID:          VerticeID(nr),
		OriginalIDs: u.CreateSet[VerticeName](),
	}
	v.OriginalIDs.Add(VerticeName(name))
	return v
}

func (v Vertice) String() string {
	return fmt.Sprintf("%d: %s", v.ID, v.OriginalIDs)
}

type Cut struct {
	Cost   uint
	Cuts   []Edge
	Group1 u.Set[VerticeName]
	Group2 u.Set[VerticeName]
}

func (c Cut) String() string {
	g1 := make([]string, 0, len(c.Group1))
	for v := range c.Group1 {
		g1 = append(g1, string(v))
	}
	g2 := make([]string, 0, len(c.Group2))
	for v := range c.Group2 {
		g2 = append(g2, string(v))
	}
	return fmt.Sprintf("Cut:\n  Cost: %d\n  Cuts: %v\n  Group1: %s\n  Group2: %s", c.Cost, c.Cuts, g1, g2)
}

func (c Cut) Clone() Cut {
	g1 := c.Group1.Clone()
	g2 := c.Group2.Clone()
	return Cut{
		Cost:   c.Cost,
		Cuts:   c.Cuts,
		Group1: g1,
		Group2: g2,
	}
}

func MinimumCut(g Graph) Cut {
	// Choose a random vertice
	var minCost uint = math.MaxUint
	var minCut Cut
	fmt.Printf("Nr vertices: %d, nr edges: %d\n", len(g.Vertices), len(g.Edges))
	for {
		nrVertices := len(g.Vertices)
		if nrVertices == 1 {
			break
		}
		// Choose a random vertice
		// Reduce the graph
		a := g.Vertices[0]
		if nrVertices > 1 {
			a = g.Vertices[1]
		}
		nrEdges := len(g.Edges)
		cut := g.Reduce(a.ID)
		fmt.Printf("Nr vertices: %d, nr edges: %d, chosen node: %d, cost: %d, minCost: %d, prod: %d\n",
			nrVertices, nrEdges, a.ID, cut.Cost, minCost, len(cut.Group1)*len(cut.Group2))
		//fmt.Println(cut)
		//fmt.Println(g)
		if cut.Cost < minCost {
			minCost = cut.Cost
			minCut = cut.Clone()
		}
	}
	return minCut
}

// Reduce reduces the number of vertices by combining two vertices s, t.
// t is the vertice with the lowest cost, and s the one with next lowest cost.
// The cost of the new vertice is the sum of the costs of the edges connecting
// t to s and the rest of the nodes.
func (g *Graph) Reduce(id VerticeID) Cut {
	usedIDs := u.CreateSet[VerticeID]()
	nrVertices := len(g.Vertices)
	usedIDsList := make([]VerticeID, 0, nrVertices)
	usedIDs.Add(id)
	usedIDsList = append(usedIDsList, id)
	left := u.CreateSet[VerticeID]()
	for _, v := range g.Vertices {
		left.Add(v.ID)
	}
	left.Remove(id)
	for {
		leftList := left.Values()
		maxCost := uint(0)
		var maxID VerticeID
		edgeMap := make(map[VerticeID]map[VerticeID]uint)
		for _, e := range g.Edges {
			_, ok := edgeMap[e.From]
			if !ok {
				edgeMap[e.From] = make(map[VerticeID]uint)
			}
			edgeMap[e.From][e.To] = e.Cost
			_, ok = edgeMap[e.To]
			if !ok {
				edgeMap[e.To] = make(map[VerticeID]uint)
			}
			edgeMap[e.To][e.From] = e.Cost
		}
		for _, lID := range leftList {
			sum := uint(0)
			fromMap := edgeMap[lID]
			for uID := range usedIDs {
				if cost, ok := fromMap[uID]; ok {
					sum += cost
				}
			}
			if sum > maxCost {
				maxCost = sum
				maxID = lID
			}
		}
		usedIDs.Add(maxID)
		usedIDsList = append(usedIDsList, maxID)
		left.Remove(maxID)
		if len(left) == 0 {
			break
		}
	}
	//fmt.Printf("Used IDs: %v\n", usedIDsList)
	if len(usedIDs) != nrVertices {
		panic("Not all vertices used")
	}
	// Add the last vertice t to the next to last vertice s
	tID := usedIDsList[nrVertices-1]
	sID := usedIDsList[nrVertices-2]
	tIdx := -1
	sIdx := -1
	var s, t Vertice
	for i, v := range g.Vertices {
		if v.ID == sID {
			s = v
			sIdx = i
		}
		if v.ID == tID {
			t = v
			tIdx = i
		}
	}
	for tID := range t.OriginalIDs {
		s.OriginalIDs.Add(tID)
	}
	g.Vertices[sIdx] = s

	// Find nodes that had edges to both s and t and update weights
	var tEdgesIdx []int
nodeLoop:
	for n_i, n := range usedIDsList[:nrVertices-1] {
		sEdgeIdx := -1
		tEdgeIdx := -1
		for i, e := range g.Edges {
			if e.From == n || e.To == n {
				if e.From == t.ID || e.To == t.ID {
					tEdgesIdx = append(tEdgesIdx, i)
					if n_i == nrVertices-2 {
						// Found s-t edge, just add to remove list
						continue nodeLoop
					}
					tEdgeIdx = i
				} else if e.From == s.ID || e.To == s.ID {
					sEdgeIdx = i
				}
			}
		}
		if sEdgeIdx != -1 && tEdgeIdx != -1 {
			g.Edges[sEdgeIdx].Cost += g.Edges[tEdgeIdx].Cost
		} else if tEdgeIdx != -1 {
			// Need to change the edge n-t to n-s
			w := g.Edges[tEdgeIdx].Cost
			e := Edge{
				From: n,
				To:   s.ID,
				Cost: w,
			}
			g.Edges = append(g.Edges, e)
		}
	}
	cut := Cut{
		Cost:   0,
		Cuts:   []Edge{},
		Group1: u.CreateSet[VerticeName](),
		Group2: u.CreateSet[VerticeName](),
	}
	// Remove all edges to t
	sort.Ints(tEdgesIdx)
	for i := len(tEdgesIdx) - 1; i >= 0; i-- {
		cut.Cost += g.Edges[tEdgesIdx[i]].Cost
		cut.Cuts = append(cut.Cuts, g.Edges[tEdgesIdx[i]])
		g.Edges = append(g.Edges[:tEdgesIdx[i]], g.Edges[tEdgesIdx[i]+1:]...)
	}
	cut.Group1 = t.OriginalIDs.Clone()
	for _, v := range g.Vertices {
		for oID := range v.OriginalIDs {
			if !cut.Group1.Contains(oID) {
				cut.Group2.Add(oID)
			}
		}
	}
	// Remove t from vertices
	g.Vertices = append(g.Vertices[:tIdx], g.Vertices[tIdx+1:]...)
	return cut
}
