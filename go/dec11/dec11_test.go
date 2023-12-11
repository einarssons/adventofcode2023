package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task(lines, 2)
	require.Equal(t, 374, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task(lines, 10)
	require.Equal(t, 1030, result)
}
