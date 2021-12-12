package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay11(t *testing.T) {
	testInput := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	p1, err := solveDay11P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 1656, p1)

	p2, err := solveDay11P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 195, p2)
}
