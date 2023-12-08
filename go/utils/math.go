package utils

import "fmt"

const (
	MaxInt = 9223372036854775807
)

func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinMaxInts(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for _, nr := range numbers {
		if nr > max {
			max = nr
		}
		if nr < min {
			min = nr
		}
	}
	return min, max
}

func Min(numbers []int) int {
	minNr := 1 << 40
	for _, nr := range numbers {
		if nr < minNr {
			minNr = nr
		}
	}
	return minNr
}

func Max(numbers []int) int {
	maxNr := -(1 << 40)
	for _, nr := range numbers {
		if nr > maxNr {
			maxNr = nr
		}
	}
	return maxNr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Sign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

func Cmp(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func Triangle(nr int) int {
	return nr * (nr + 1) / 2
}

// GCD - greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

// LCM - least common multiple (LCM) via GCD
func LCM(values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		result = result * values[i] / GCD(result, values[i])
	}
	return result
}

// Cycle is a struct that represents a cycle with an offset and a period.
type Cycle struct {
	Offset, Period int
}

// CRT is a function that takes a slice of cycles and returns the smallest
// number that satisfies all the cycles according to the Chinese Remainder Theorem.
func CRT(cycles []Cycle) int {
	n := 1
	for _, c := range cycles {
		d := GCD(n, c.Period)
		n *= c.Period / d
	}

	for i := 0; i < len(cycles); i++ {
		term := 0
		prod := 1
		for j := 0; j < len(cycles); j++ {
			if i != j {
				prod *= cycles[j].Period
			}
		}
		found := false
		for k := 1; k < cycles[i].Period; k++ {
			if cycles[i].Offset == 0 {
				found = true
				break
			}
			term = k * prod
			if term%cycles[i].Period == cycles[i].Offset {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("no extra sum found for cycle", i)
			return -1
		}
		n += term
	}
	return n
}
