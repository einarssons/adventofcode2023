package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, 94, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task2(lines)
	require.Equal(t, 0, result)
}
