package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	x, y := 10, 6
	l := NewGame(x, y)

	assert.Equal(t, y, len(l.Grid))
	assert.Equal(t, 0, l.Grid[0][0])
	assert.Equal(t, 0, l.Iterations)
}

func TestRandomizeGame(t *testing.T) {
	x, y := 10, 10
	l := NewGame(x, y)
	g := &l.Grid
	l.Randomize()

	assert.NotEqual(t, g, l.Grid)
}
