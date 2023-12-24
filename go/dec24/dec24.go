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
		fmt.Println("task1: ", task1(lines, 200000000000000, 400000000000000))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string, minXY, maxXY float64) int {
	hails := parse(lines)
	nrCrosses := 0
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			cross := pathCrossesXY(hails[i], hails[j], minXY, maxXY)
			if cross {
				nrCrosses++
			}
		}
	}
	return nrCrosses
}

func task2(lines []string) int {
	hails := parse(lines)
	findParallel(hails)
	throw := perfectThrow(hails)
	return throw.px + throw.py + throw.pz
}

func findParallel(hails []hailstone) {
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			parallel := isParallel(hails[i], hails[j])
			if parallel {
				fmt.Printf("%v and %v are parallel\n", hails[i], hails[j])
			}
		}
	}
}

func isParallel(h1, h2 hailstone) bool {
	// h1.vx / h2.vx = h1.vy / h2.vy = h1.vz / h2.vz
	if h1.vx*h2.vy == h2.vx*h1.vy && h1.vx*h2.vz == h2.vx*h1.vz {
		return true
	}
	return false
}

func pathCrossesXY(h1, h2 hailstone, minXY, maxXY float64) bool {
	/*
		// h1.px + h1.vx * t1 = h2.px + h2.vx * t2
		// h1.py + h1.vy * t1 = h2.py + h2.vy * t2
		// h1.vx * t1 - h2.vx * t2 = h2.px - h1.px
		// h1.vy * t1 - h2.vy * t2 = h2.py - h1.py
		// solve for t1 and t2
		// h2.vy*t2 = h1.vy*t1 - h2.py + h1.py
		// t2 = (h1.vy*t1 - h2.py + h1.py) / h2.vy
		// h1.vx * t1 - h2.vx * (h1.vy*t1 - h2.py + h1.py) / h2.vy = h2.px - h1.px
		// h1.vx * h2.vy * t1 - h2.vx * (h1.vy*t1 - h2.py + h1.py) = (h2.px - h1.px) * h2.vy
		// h1.vx * h2.vy * t1 - h2.vx * h1.vy*t1 + h2.vx * h2.py - h2.vx * h1.py = (h2.px - h1.px) * h2.vy
		// t1 * (h1.vx * h2.vy - h2.vx * h1.vy) = (h2.px - h1.px) * h2.vy - h2.vx * h2.py + h2.vx * h1.py
		// t1 = ((h2.px - h1.px) * h2.vy - h2.vx * h2.py + h2.vx * h1.py) / (h1.vx * h2.vy - h2.vx * h1.vy)
	*/
	if h1.vx*h2.vy == h2.vx*h1.vy {
		return false
	}
	h1px := float64(h1.px)
	h1py := float64(h1.py)
	h1vx := float64(h1.vx)
	h1vy := float64(h1.vy)
	h2px := float64(h2.px)
	h2py := float64(h2.py)
	h2vx := float64(h2.vx)
	h2vy := float64(h2.vy)
	t1 := ((h2px-h1px)*h2vy - h2vx*h2py + h2vx*h1py) / (h1vx*h2vy - h2vx*h1vy)
	if t1 < 0 {
		return false
	}
	t2 := (h1vx*t1 - h2px + h1px) / h2vx
	if t2 < 0 {
		return false
	}
	px := h1px + h1vx*t1
	py := h1py + h1vy*t1
	if px >= minXY && px <= maxXY && py >= minXY && py <= maxXY {
		//fmt.Printf("%v %v cross at %.3f at %.3f,%.3f\n", h1, h2, t1, px, py)
		return true
	}
	return false
}

type hailstone struct {
	px, py, pz, vx, vy, vz int
}

func parse(lines []string) []hailstone {
	stones := make([]hailstone, 0, len(lines))
	for _, line := range lines {
		p1, p2 := u.Cut(line, "@")
		c1 := strings.Split(p1, ",")
		c2 := strings.Split(p2, ",")
		s := hailstone{u.Atoi(c1[0]), u.Atoi(c1[1]), u.Atoi(c1[2]), u.Atoi(c2[0]), u.Atoi(c2[1]), u.Atoi(c2[2])}
		stones = append(stones, s)
	}
	return stones
}

// perfectThrow finds the throw that hits all hailstones at some
// time in the future, where each time is different.
func perfectThrow(hails []hailstone) (throw hailstone) {
	// r(t) = (px + vx * t, py + vy * t, pz + vz * t)
	// r_i(t) = (px_i + vx_i * t, py_i + vy_i * t, pz_i + vz_i * t)
	// r(t_0) = r_0(t_0) // 3 equations, 7 unknowns
	// r(t_1) = r_1(t_1) // 6 equations, 8 unknowns
	// r(t_2) = r_2(t_2) // 9 equations, 9 unknowns
	// px + vx * t_0 = px_0 + vx_0 * t_0
	// py + vy * t_0 = py_0 + vy_0 * t_0
	// pz + vz * t_0 = pz_0 + vz_0 * t_0
	// px + vx * t_1 = px_1 + vx_1 * t_1
	// py + vy * t_1 = py_1 + vy_1 * t_1
	// pz + vz * t_1 = pz_1 + vz_1 * t_1
	// px + vx * t_2 = px_2 + vx_2 * t_2
	// py + vy * t_2 = py_2 + vy_2 * t_2
	// pz + vz * t_2 = pz_2 + vz_2 * t_2

	fmt.Println("solve this using Z3 on the the first three hailstones")

	throw = hailstone{24, 13, 10, -3, 1, 2}
	return throw
}
