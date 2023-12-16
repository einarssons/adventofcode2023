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
	l := strings.Join(lines, "")
	parts := strings.Split(l, ",")
	sum := 0
	for _, p := range parts {
		sum += hash(p)
	}
	return sum
}

func hash(s string) int {
	runes := u.SplitToRunes(s)
	hash := rune(0)
	for _, r := range runes {
		hash += r
		hash *= 17
		hash = hash % 256
	}
	return int(hash)
}

func task2(lines []string) int {
	l := strings.Join(lines, "")
	steps := strings.Split(l, ",")
	boxes := make([]box, 256)
	for _, s := range steps {
		processStep(s, boxes)
	}
	return calcFocalLength(boxes)
}

func processStep(s string, boxes []box) {
	switch {
	case strings.HasSuffix(s, "-"):
		lbl := strings.TrimSuffix(s, "-")
		removeLens(lbl, boxes)
	default:
		lbl, val := u.Cut(s, "=")
		fLen := u.Atoi(val)
		addLens(lbl, fLen, boxes)
	}
}

func addLens(lbl string, fLen int, boxes []box) {
	i := hash(lbl)
	for j, l := range boxes[i].lenses {
		if l.lbl == lbl {
			boxes[i].lenses[j].fLen = fLen
			return
		}
	}
	boxes[i].lenses = append(boxes[i].lenses, lens{lbl, fLen})
}

func removeLens(lbl string, boxes []box) {
	i := hash(lbl)
	for j, l := range boxes[i].lenses {
		if l.lbl == lbl {
			boxes[i].lenses = append(boxes[i].lenses[:j], boxes[i].lenses[j+1:]...)
			return
		}
	}
}

func calcFocalLength(boxes []box) int {
	sum := 0
	for i, b := range boxes {
		for j, l := range b.lenses {
			sum += (i + 1) * (j + 1) * l.fLen
		}
	}
	return sum
}

type lens struct {
	lbl  string
	fLen int
}

type box struct {
	lenses []lens
}
