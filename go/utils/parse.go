package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/oriser/regroup"
)

// SplitToInts finds all ints in line (including sign).
func SplitToInts(line string) []int {
	re := regexp.MustCompile("-?[0-9]+")
	parts := re.FindAllString(line, -1)
	var numbers []int
	for _, p := range parts {
		number, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

// SplitToChars splits a line into chars.
func SplitToChars(line string) []string {
	chars := make([]string, len(line))
	for i := 0; i < len(line); i++ {
		chars[i] = line[i : i+1]
	}
	return chars
}

// SplitToRunes splits a line into runes.
func SplitToRunes(line string) []rune {
	runes := make([]rune, 0, len(line))
	for i := 0; i < len(line); i++ {
		runes = append(runes, rune(line[i]))
	}
	return runes
}

func ContainsInt(x int, entries []int) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

// Trim just trims space.
func Trim(line string) string {
	return strings.TrimSpace(line)
}

func SplitWithTrim(line, splitPattern string) []string {
	if splitPattern == "" || splitPattern == " " {
		panic("splitPattern cannot be empty or space")
	}
	parts := strings.Split(line, splitPattern)
	for i, p := range parts {
		parts[i] = Trim(p)
	}
	return parts
}

// SplitWithSpace splits a line with space as separator.
func SplitWithSpace(line string) []string {
	return strings.Split(line, " ")
}

func CountInts(entries []int) map[int]int {
	m := make(map[int]int)
	for _, n := range entries {
		m[n]++
	}
	return m
}

func CountStrings(entries []string) map[string]int {
	m := make(map[string]int)
	for _, n := range entries {
		m[n]++
	}
	return m
}

func ContainsString(x string, entries []string) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

var rex = regroup.MustCompile(`(?P<verb>[a-zA-Z]+)\s+(?P<value>\d+)`)

type Command struct {
	Verb  string `regroup:"verb"`
	Value int    `regroup:"value"`
}

// ParseCommand parses a "verb value" from a line.
func ParseCommand(line string) Command {
	c := Command{}
	if err := rex.MatchToTarget(line, &c); err != nil {
		log.Fatal(err)
	}
	return c
}

// FirstAsciiNr return ascii number of first character in string.
// Panics if value > 127.
func FirstAsciiNr(s string) int {
	runeValue := int([]rune(s)[0])
	if runeValue > 127 {
		panic(fmt.Sprintf("rune value %d is not ascii", runeValue))
	}
	return runeValue
}

// Atoi is panicing version of strconv.Atoi. Panics instead of error.
func Atoi(s string) int {
	s = strings.TrimSpace(s)
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", s))
	}
	return nr
}

// Cut is panicing version of strings.Cut. Panics instead of error.
func Cut(s string, separator string) (left, right string) {
	var ok bool
	left, right, ok = strings.Cut(s, separator)
	if !ok {
		panic(fmt.Sprintf("cannot cut %q with %q", s, separator))
	}
	return left, right
}
