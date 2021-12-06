package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay6(t *testing.T) {
	testInput := `3,4,3,1,2`

	p218, err := solveDay6(testInput, 18)
	require.NoError(t, err)
	require.Equal(t, 26, p218)

	p280, err := solveDay6(testInput, 80)
	require.NoError(t, err)
	require.Equal(t, 5934, p280)

	p2256, err := solveDay6(testInput, 256)
	require.NoError(t, err)
	require.Equal(t, 26984457539, p2256)
}
