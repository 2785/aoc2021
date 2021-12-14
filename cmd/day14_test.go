package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay14(t *testing.T) {
	testInput := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	p1, err := solveDay14P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 1588, p1)

	p2, err := solveDay14P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 2188189693529, p2)
}
