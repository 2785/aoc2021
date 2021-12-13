package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay13(t *testing.T) {
	testInput := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

	p1, err := solveDay13P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 17, p1)
}
