package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2023/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines, 6)
	require.Equal(t, 16, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	testCases := []struct {
		maxSteps      int
		expectedPlots int
	}{
		{maxSteps: 6, expectedPlots: 16},
		{maxSteps: 10, expectedPlots: 50},
		{maxSteps: 50, expectedPlots: 1594},
		{maxSteps: 100, expectedPlots: 6536},
		{maxSteps: 500, expectedPlots: 167004},
		{maxSteps: 1000, expectedPlots: 668697},
		{maxSteps: 5000, expectedPlots: 16733044},
	}
	for _, tc := range testCases {
		result := task1(lines, tc.maxSteps)
		require.Equal(t, tc.expectedPlots, result)
	}
}
