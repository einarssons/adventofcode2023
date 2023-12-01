package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadNumbersFromFile reads one integer from each line in file.
func ReadNumbersFromFile(path string) []int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var numbers []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := s.Text()
		nr, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, nr)
	}
	return numbers
}

// ReadLinesFromFile reads lines from a file.
func ReadLinesFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		trimmed := strings.Trim(line, " ")
		lines = append(lines, trimmed)
	}
	return lines
}

// TrimTrailingNewline removes trailing newline
func TrimTrailingNewline(lines []string) []string {
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

// ReadRawLinesFromFile reads lines from a file.
func ReadRawLinesFromFile(path string) []string {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	rawStr := string(raw)
	return strings.Split(rawStr, "\n")
}
