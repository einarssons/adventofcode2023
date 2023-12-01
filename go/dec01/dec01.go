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
			digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			for j, digit := range digits {
				if i+len(digit) <= l && line[i:i+len(digit)] == digit {
					d = j + 1
				}
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
