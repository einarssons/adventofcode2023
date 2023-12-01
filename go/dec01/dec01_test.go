package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, 142, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput2")
	result := task2(lines)
	require.Equal(t, 281, result)
}
