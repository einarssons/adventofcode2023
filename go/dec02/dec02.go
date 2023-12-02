package main

import (
	"flag"
	"fmt"

	u "github.com/einarssons/adventofcode2023/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines, 12, 13, 14))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string, red, green, blue int) int {
	sum := 0
	for _, line := range lines {
		g := parseLine(line)
		ok := true
		for _, d := range g.draws {
			if d.red > red || d.green > green || d.blue > blue {
				ok = false
				break
			}
		}
		if ok {
			sum += g.nr
		}
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for _, line := range lines {
		g := parseLine(line)
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, d := range g.draws {
			if d.red > minRed {
				minRed = d.red
			}
			if d.green > minGreen {
				minGreen = d.green
			}
			if d.blue > minBlue {
				minBlue = d.blue
			}
		}
		power := minRed * minGreen * minBlue
		sum += power
	}
	return sum
}

type game struct {
	nr    int
	draws []draw
}

type draw struct {
	red   int
	green int
	blue  int
}

func parseLine(line string) game {
	parts := u.SplitWithTrim(line, ":")
	p2 := u.SplitWithSpace(parts[0])
	nr := u.Atoi(p2[1])
	g := game{nr: nr}
	draws := u.SplitWithTrim(parts[1], ";")
	for _, list := range draws {
		d := draw{}
		parts := u.SplitWithTrim(list, ",")
		for _, p := range parts {
			p2 := u.SplitWithSpace(p)
			count := u.Atoi(p2[0])
			col := p2[1]
			switch col {
			case "red":
				d.red = count
			case "green":
				d.green = count
			case "blue":
				d.blue = count
			default:
				panic("unknown color")
			}
		}
		g.draws = append(g.draws, d)
	}
	return g
}
