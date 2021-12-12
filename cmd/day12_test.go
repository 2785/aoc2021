package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay12(t *testing.T) {
	testInput1 := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	testInput2 := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

	p1t1, err := solveDay12P1(testInput1)
	require.NoError(t, err)
	require.Equal(t, 10, p1t1)

	p1t2, err := solveDay12P1(testInput2)
	require.NoError(t, err)
	require.Equal(t, 19, p1t2)

	p2t1, err := solveDay12P2(testInput1)
	require.NoError(t, err)
	require.Equal(t, 36, p2t1)
}
