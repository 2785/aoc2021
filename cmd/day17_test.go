package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay17(t *testing.T) {
	testInput := "target area: x=20..30, y=-10..-5"
	p1, err := solveDay17P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 45, p1)

	p2, err := solveDay17P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 112, p2)
}
