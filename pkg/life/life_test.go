package life

import (
	"fmt"
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

func TestNewGameFromSeed(t *testing.T) {
	g := [][][]int{
		{
			{0, 1, 1},
			{1, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 0, 0},
			{0, 1},
			{0, 0, 0},
		},
		{
			{0, 0, 0},
			{0, 0, 2},
			{0, 0, 1},
		},
	}
	g1, err := NewGameFromSeed(g[0])
	assert.Equal(t, g[0], g1.Grid)
	assert.Nil(t, err)
	g2, err := NewGameFromSeed(g[1])
	assert.Nil(t, g2)
	assert.Error(t, err)
	g3, err := NewGameFromSeed(g[2])
	assert.Nil(t, g3)
	assert.Error(t, err)
}

func TestRandomizeGame(t *testing.T) {
	// check if different from NewGame
	x, y := 15, 4
	l := NewGame(x, y)
	g := &l.Grid
	l.Randomize()

	// check for duplicate rows
	assert.NotEqual(t, g, l.Grid)
	for i, j := range l.Grid {
		for m, n := range l.Grid {
			if i == m {
				continue
			}
			assert.NotEqual(t, j, m, fmt.Sprintf("Duplicate rows are present: indexes %d and %d", i, n))
		}
	}

	// check if can be repeated
	g = &l.Grid
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
	g.Grid = [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}
	assert.Equal(t, 3, g.countNeighbors([2]int{2, 1}))
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

	g.Grid = [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	expectedCells = [][2]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}
	assert.Equal(t, expectedCells, g.cellsToCheck())
}

func TestIfCellLives(t *testing.T) {
	g := &Game{Grid: [][]int{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}}
	assert.Equal(t, 1, g.ifCellLives([2]int{1, 1}))
	assert.Equal(t, 1, g.ifCellLives([2]int{1, 2}))
	assert.Equal(t, 0, g.ifCellLives([2]int{2, 1}))
}

func TestIterate(t *testing.T) {
	g := &Game{[][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}, 0}

	var iterations = [][][]int{
		{
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}
	r := 0

	for _, i := range iterations {
		g.Iterate()
		r++
		assert.Equal(t, i, g.Grid)
		assert.Equal(t, r, g.Iterations)
	}
}
