package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay10(t *testing.T) {

	testInput := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	p1, err := SolveDay10P1(testInput)
	require.NoError(t, err)
	require.Equal(t, 26397, p1)

	p2, err := SolveDay10P2(testInput)
	require.NoError(t, err)
	require.Equal(t, 288957, p2)

}
