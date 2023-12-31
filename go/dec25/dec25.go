package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines))
	}
}

func stopProfile(ctx context.Context) {
	<-ctx.Done()
	pprof.StopCPUProfile()
	fmt.Println("wrote cpu.prof")
}

func task1(lines []string) int {
	graph := parse(lines)
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 20*time.Second)
	go stopProfile(ctx)
	defer pprof.StopCPUProfile()
	cut := MinimumCut(graph)
	fmt.Println("Minimal cut: ", cut)
	return len(cut.Group1) * len(cut.Group2)
}

func parse(lines []string) Graph {
	fh, err := os.Create("nodes.csv")
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	m := u.CreateSet[string]()
	g := Graph{}
	knownVertices := u.CreateSet[string]()
	for _, line := range lines {
		k, vs := u.Cut(line, ":")
		vs = strings.TrimSpace(vs)
		values := strings.Split(vs, " ")
		if !knownVertices.Contains(k) {
			g.Vertices = append(g.Vertices, NewVertice(len(g.Vertices)+1, k))
			knownVertices.Add(k)
		}
		for _, v := range values {
			kv := k + "-" + v
			m.Add(kv)
			fmt.Fprintf(fh, "%s,%s\n", k, v)
			if !knownVertices.Contains(v) {
				g.Vertices = append(g.Vertices, NewVertice(len(g.Vertices)+1, v))
				knownVertices.Add(v)
			}
			v1 := g.VerticeIDFromName(VerticeName(k))
			v2 := g.VerticeIDFromName(VerticeName(v))
			g.Edges = append(g.Edges, Edge{v1, v2, 1})
		}
	}

	return g
}
