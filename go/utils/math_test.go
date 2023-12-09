package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRT(t *testing.T) {
	cycles := []Cycle{
		{0, 3},
		{3, 4},
		{4, 5},
	}
	result := CRT(cycles)
	require.Equal(t, 39, result)
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 1},
		{3, 12, 3},
		{24, 25, 1},
	}

	for _, tc := range testCases {
		got := GCD(tc.a, tc.b)
		require.Equal(t, tc.expected, got, fmt.Sprintf("GCD(%d,%d)\n", tc.a, tc.b))
	}
}
