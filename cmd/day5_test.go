package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay5(t *testing.T) {
	testInput := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	p1, err := solveDay5P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 5, p1)

	p2, err := solveDay5P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 12, p2)
}
