package life

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// In Grid cells, 0 is dead and 1 is alive
type Game struct {
	Grid       [][]int
	Iterations int
}

// Helper func for cellsToCheck private method
func appendIfMissing(slice [][2]int, i [2]int) [][2]int {
	for _, e := range slice {
		if e == i {
			return slice
		}
	}
	return append(slice, i)
}

// Counts live cells around a cell
func (g *Game) countNeighbors(location [2]int) int {
	// coordinates must be in an array of [Row, Column] format
	grid := g.Grid
	y, x := location[0], location[1]
	lenY, lenX := len(grid), len(grid[0])
	var neighbors int

	for i := -1; i < 2; i++ {
		if (i+y > -1) && (i+y < lenY) {
			for j := -1; j < 2; j++ {
				if j == 0 && i == 0 {
					continue
				}
				if j+x > -1 && j+x < lenX {
					neighbors += grid[i+y][j+x]
				}
			}
		}
	}
	return neighbors
}

// Find all cells touching alive ones.
// Used during lifecycle iteration to find which cells to check for life.
func (g *Game) cellsToCheck() [][2]int {
	lenY, lenX := len(g.Grid), len(g.Grid[0])
	var cells [][2]int

	// i,j for grid y,x
	// m,n for neighbors y,x
	for i := range g.Grid {
		for j := range g.Grid[i] {
			if g.Grid[i][j] == 1 {
				for m := -1; m < 2; m++ {
					if i+m > -1 && i+m < lenY {
						for n := -1; n < 2; n++ {
							if j+n > -1 && j+n < lenX {
								cells = appendIfMissing(cells, [2]int{i + m, j + n})
							}
						}
					}
				}
			}
		}
	}
	return cells
}

// Takes in [y,x] as cell index.
// Returns 1 if cell lives, 0 if it dies.
func (g *Game) ifCellLives(location [2]int) int {
	cell := g.Grid[location[0]][location[1]]
	neighbors := g.countNeighbors(location)
	total := cell + neighbors
	if neighbors == 3 || total == 3 {
		return 1
	}
	return 0
}

// Game Methods

// Iterates the Game through one generation
func (g *Game) Iterate() {
	changes := make(map[[2]int]int)
	c := g.cellsToCheck()
	for _, cell := range c {
		status := g.ifCellLives(cell)
		if status != g.Grid[cell[0]][cell[1]] {
			changes[cell] = status
		}
	}
	for cell, val := range changes {
		g.Grid[cell[0]][cell[1]] = val
	}
	g.Iterations++
}

// Utility method that prints the Grid to terminal as lines of strings
func (g *Game) Print() {
	for _, i := range g.Grid {
		var r []string
		for _, j := range i {
			r = append(r, strconv.Itoa(j))
		}
		fmt.Println(strings.Join(r, ""))
	}
}

// Sets all Game.Grid cell values to randomly either 0 or 1
func (g *Game) Randomize() {
	rand.Seed(time.Now().UnixNano())
	for i := range g.Grid {
		for j := range g.Grid[i] {
			g.Grid[i][j] = rand.Intn(2)
		}
	}
}

// Returns a new Game struct with y rows and x cells per row.
func NewGame(x int, y int) *Game {
	grid := make([][]int, y)
	for i := range grid {
		grid[i] = make([]int, x)
	}
	return &Game{Grid: grid}
}

// Returns a new Game struct from maa 2D array of 0s and 1s.
func NewGameFromSeed(seed [][]int) (*Game, error) {
	// check if all rows have same length
	var columnCounts []int
	for _, i := range seed {
		columnCounts = append(columnCounts, len(i))
		for _, j := range i {
			if j < 0 || j > 1 {
				return nil, errors.New("Proveded seed array contains a cell value that is not 0 or 1")
			}
		}
	}
	l := columnCounts[0]
	for _, i := range columnCounts[1:] {
		if i != l {
			return nil, errors.New("Provided seed array does not have rows of all the same length.")
		}
	}
	return &Game{Grid: seed}, nil
}
