package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	input := `199
200
208
210
200
207
240
269
260
263`

	out, err := SolveDay1P1(input)
	assert.NoError(t, err)
	assert.Equal(t, 7, out)

	out2, err := SolveDay1P2(input)
	assert.NoError(t, err)
	assert.Equal(t, 5, out2)
}
