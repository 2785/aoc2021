package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay7(t *testing.T) {

	testInput := `16,1,2,0,4,2,7,1,2,14`

	p1, err := solveDay7P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 37, p1)

	p2, err := solveDay7P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 168, p2)

}
