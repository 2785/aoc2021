package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay9(t *testing.T) {
	testInput := `2199943210
3987894922
9856789892
8767896789
9899965678`

	p1, err := solveDay9P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 15, p1)

	p2, err := solveDay9P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 1134, p2)

	// 	otherTestInput := `123
	// 894
	// 765`

	// 	p2, err := solveDay9P2(otherTestInput)
	// 	require.NoError(t, err)
	// 	require.Equal(t, 8, p2)
}
