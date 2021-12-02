package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {

	data := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	p1, err := solveDay2P1(data)
	assert.NoError(t, err)
	assert.Equal(t, 150, p1)

	p2, err := solveDay2P2(data)
	assert.NoError(t, err)
	assert.Equal(t, 900, p2)
}
