package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay21(t *testing.T) {
	testInput := `Player 1 starting position: 4
Player 2 starting position: 8`

	p1, err := solveDay21P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 739785, p1)

	p2, err := solveDay21P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 444356092776315, p2)
}
