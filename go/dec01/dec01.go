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
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		firstDigit := -1
		lastDigit := -1
		for i := 0; i < len(line); i++ {
			c := int(line[i])
			d := -1
			if c >= 48 && c <= 57 {
				d = c - 48
			}
			if d >= 0 {
				if firstDigit < 0 {
					firstDigit = d
				}
				lastDigit = d
			}

		}
		nr := firstDigit*10 + lastDigit
		sum += nr
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for _, line := range lines {
		l := len(line)
		firstDigit := -1
		lastDigit := -1
		for i := 0; i < len(line); i++ {
			c := int(line[i])
			d := -1
			if c >= 48 && c <= 57 {
				d = c - 48
			}
			switch {
			case i+4 <= l && line[i:i+4] == "zero":
				d = 0
			case i+3 <= l && line[i:i+3] == "one":
				d = 1
			case i+3 <= l && line[i:i+3] == "two":
				d = 2
			case i+5 <= l && line[i:i+5] == "three":
				d = 3
			case i+4 <= l && line[i:i+4] == "four":
				d = 4
			case i+4 <= l && line[i:i+4] == "five":
				d = 5
			case i+3 <= l && line[i:i+3] == "six":
				d = 6
			case i+5 <= l && line[i:i+5] == "seven":
				d = 7
			case i+5 <= l && line[i:i+5] == "eight":
				d = 8
			case i+4 <= l && line[i:i+4] == "nine":
				d = 9
			}
			if d >= 0 {
				if firstDigit < 0 {
					firstDigit = d
				}
				lastDigit = d
			}

		}
		nr := firstDigit*10 + lastDigit
		sum += nr
	}
	return sum
}
