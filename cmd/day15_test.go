package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay15(t *testing.T) {
	testInput := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	p1, err := solveDay15P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 40, p1)

	p2, err := solveDay15P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 315, p2)
}
