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
	g := &Game{[][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}, 0}
	assert.Equal(t, 0, g.countNeighbors([2]int{1, 1}))

	g.Grid = [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	assert.Equal(t, 1, g.countNeighbors([2]int{1, 1}))

	g.Grid = [][]int{
		{1, 1, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	assert.Equal(t, 3, g.countNeighbors([2]int{1, 1}))

	g.Grid = [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	assert.Equal(t, 2, g.countNeighbors([2]int{2, 2}))
}

func TestFindAliveCells(t *testing.T) {
	g := &Game{[][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}, 0}

	expectedCells := [][2]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}

	assert.Equal(t, expectedCells, g.cellsToCheck())
}
