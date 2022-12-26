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

func TestCountNeighbors(t *testing.T) {
	testGrid1 := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 0, countNeighbors(testGrid1, [2]int{1, 1}))

	testGrid2 := [][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 1, countNeighbors(testGrid2, [2]int{1, 1}))

	testGrid3 := [][]int{
		[]int{1, 1, 1},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 3, countNeighbors(testGrid3, [2]int{1, 1}))

	testGrid4 := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 2, countNeighbors(testGrid4, [2]int{2, 2}))
}
