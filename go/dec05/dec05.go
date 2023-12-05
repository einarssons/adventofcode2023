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
	seeds, maps := parse(lines)
	minDist := u.MaxInt
	for _, s := range seeds {
		val := s
		for _, m := range maps {
			for _, r := range m.ranges {
				if val >= r.srcStart && val < r.srcStart+r.len {
					val = r.dstStart + (val - r.srcStart)
					break
				}
			}
		}
		if val < minDist {
			minDist = val
		}
	}
	return minDist
}

func task2(lines []string) int {
	itvls, maps := parse2(lines)
	minDist := u.MaxInt
	for i, itvl := range itvls {
		fmt.Printf("interval %d\n", i)
		for v := itvl.start; v < itvl.end(); v++ {
			val := v
			for _, m := range maps {
				for _, r := range m.ranges {
					if val >= r.srcStart && val < r.srcStart+r.len {
						val = r.dstStart + (val - r.srcStart)
						break
					}
				}
			}
			if val < minDist {
				minDist = val
			}
		}
	}
	return minDist
}

type itvl struct {
	dstStart, srcStart, len int
}

type Map struct {
	name   string
	ranges []itvl
}

func parse(lines []string) (seeds []int, maps []*Map) {
	newMap := false
	for i, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case i == 0:
			seeds = u.SplitToInts(line[7:])
		case line == "":
			newMap = true
		case newMap:
			newMap = false
			parts := strings.Fields(line)
			m := Map{
				name: parts[0],
			}
			maps = append(maps, &m)
		default:
			nrs := u.SplitToInts(line)
			m := maps[len(maps)-1]
			r := itvl{
				dstStart: nrs[0],
				srcStart: nrs[1],
				len:      nrs[2],
			}
			m.ranges = append(m.ranges, r)
		}
	}
	return seeds, maps
}

type Range struct {
	start, len int
}

func (r Range) end() int {
	return r.start + r.len
}

func parse2(lines []string) (seedItvls []Range, maps []*Map) {
	newMap := false
	for i, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case i == 0:
			seeds := u.SplitToInts(line[7:])
			seedItvls = make([]Range, len(seeds)/2)
			for i := 0; i < len(seeds); i += 2 {
				seedItvls[i/2] = Range{
					start: seeds[i],
					len:   seeds[i+1],
				}
			}
		case line == "":
			newMap = true
		case newMap:
			newMap = false
			parts := strings.Fields(line)
			m := Map{
				name: parts[0],
			}
			maps = append(maps, &m)
		default:
			nrs := u.SplitToInts(line)
			m := maps[len(maps)-1]
			r := itvl{
				dstStart: nrs[0],
				srcStart: nrs[1],
				len:      nrs[2],
			}
			m.ranges = append(m.ranges, r)
		}
	}
	return seedItvls, maps
}
