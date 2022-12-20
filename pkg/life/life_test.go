package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLifeGame(t *testing.T) {
	x, y := 10, 6
	l := NewGame(x, y)

	assert.Equal(t, y, len(l.Grid))
	assert.Equal(t, 0, l.Iterations)
}
