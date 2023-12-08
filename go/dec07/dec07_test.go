package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	entries := make([]entry, 0, len(lines))
	for _, line := range lines {
		hand, bid := parse(line)
		entries = append(entries, entry{hand, bid})
	}
	for _, e := range entries {
		fmt.Printf("%v %v %d\n", e.hand, e.bid, e.hand.Type())
	}
	result := task1(lines)
	require.Equal(t, 6440, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task2(lines)
	require.Equal(t, 0, result)
}
