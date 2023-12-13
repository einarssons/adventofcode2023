package main

import (
	"flag"
	"fmt"
	"strings"

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
	rows := parse(lines)
	fmt.Printf("%+v\n", rows)
	sum := 0
	for _, r := range rows {
		m := initMemoize()
		nr := m.counts(r.pattern, r.groups, 0, false)
		fmt.Println(nr, "combinations for", r.pattern, "memSize", len(m.m))
		sum += nr
	}
	return sum
}

var m memoizer

func task2(lines []string) int {
	rows := parse(lines)
	sum := 0
	for i, r := range rows {
		m := initMemoize()
		row := extendRow(r)
		nr := m.counts(row.pattern, row.groups, 0, false)
		fmt.Printf("%4d %d combinations memSize=%d\n", i, nr, len(m.m))
		sum += nr
	}
	return sum
}

type memParams struct {
	right      string
	groups     string
	nrLeftH    int
	forceEmpty bool
}

type memoizer struct {
	m map[memParams]int
}

func (m memoizer) counts(right string, groups []int, nrLeftH int, forceEmpty bool) int {
	params := memParams{right, fmt.Sprintf("%v", groups), nrLeftH, forceEmpty}
	if v, ok := m.m[params]; ok {
		return v
	}
	v := m.countsInternal(right, groups, nrLeftH, forceEmpty)
	m.m[params] = v
	return v
}

func initMemoize() memoizer {
	return memoizer{make(map[memParams]int)}
}

func (m memoizer) countsInternal(right string, groups []int, nrLeftH int, forceEmpty bool) int {
	remainingH := countHashes(right) + nrLeftH
	if len(groups) == 0 {
		if remainingH == 0 {
			return 1
		}
		return 0
	}
	if right == "" && len(groups) > 0 {
		if len(groups) == 1 && groups[0] == nrLeftH {
			return 1
		}
		return 0
	}

	/*
		if right == "" {
			return 0
		} */
	first := right[0]
	rest := right[1:]
	if forceEmpty && first == '#' {
		return 0
	}

	if nrLeftH == groups[0] {
		return m.counts(right, groups[1:], 0, true)
	}

	if nrLeftH > 0 && nrLeftH < groups[0] && first == '.' {
		return 0 // Break group to early
	}

	if first == '?' {
		if forceEmpty {
			return m.counts("."+rest, groups, nrLeftH, false)
		}
		nr1 := m.counts("#"+rest, groups, nrLeftH, false)
		nr2 := m.counts("."+rest, groups, nrLeftH, false)
		return nr1 + nr2
	}

	if first == '#' {
		return m.counts(rest, groups, nrLeftH+1, false)
	}

	if first == '.' {
		return m.counts(rest, groups, 0, false)
	}
	panic("end of func")
}

func countHashes(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			sum += 1
		}
	}
	return sum
}

func extendRow(r row) row {
	ps := []string{}
	g := make([]int, 0, len(r.groups)*5)
	for i := 0; i < 5; i++ {
		ps = append(ps, r.pattern)
		g = append(g, r.groups...)
	}
	return row{strings.Join(ps, "?"), g}
}

type row struct {
	pattern string
	groups  []int
}

func (r row) String() string {
	return fmt.Sprintf("%s %v", r.pattern, r.groups)
}

func parse(lines []string) []row {
	rows := make([]row, len(lines))
	for i, line := range lines {
		rows[i] = parseRow(line)
	}
	return rows
}

func parseRow(line string) row {
	pattern, g, ok := strings.Cut(line, " ")
	if !ok {
		panic("invalid input")
	}
	g2 := strings.Split(g, ",")
	groups := make([]int, 0, len(g2))
	for _, r := range g2 {
		groups = append(groups, u.Atoi(r))
	}
	return row{pattern, groups}
}
