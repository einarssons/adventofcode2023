package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolGrid(t *testing.T) {
	g := CreateGrid2D[bool](4, 2)
	g.Set(true, 0, 1)
	require.Equal(t, true, g.Get(0, 1))
	r, c, ok := g.Find(true)
	require.Equal(t, 0, r)
	require.Equal(t, 1, c)
	require.Equal(t, true, ok)
}
