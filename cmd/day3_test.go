package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay3Part1(t *testing.T) {
	testInput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	p1, err := solveDay3P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 198, p1)

	p2, err := solveDay3P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 230, p2)
}
