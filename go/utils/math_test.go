package utils

import (
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
