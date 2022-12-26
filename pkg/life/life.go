package life

import (
	"fmt"
	"math/rand"
	"time"
)

// In Grid cells, 0 is dead and 1 is alive
type Game struct {
	Grid       [][]int
	Iterations int
}

// Counts live cells around a cell
func (g *Game) countNeighbors(location [2]int) int {
	// coordinates must be in an array of [Row, Column] format
	grid := g.Grid
	y, x := location[0], location[1]
	lenY, lenX := len(grid), len(grid[0])
	var neighbors int

	for i := -1; i < 2; i++ {
		if (i+y != -1) && (i+y != lenY) {
			for j := -1; j < 2; j++ {
				if j == 0 && i == 0 {
					continue
				}
				if j+x != -1 && j+x != lenX {
					neighbors += grid[i+y][j+x]
				}
			}
		}
	}
	return neighbors
}

// Find all cells touching alive ones
// func (g *Game) findAliveCells() [][]int {
// 	var cells [][]int
// 	for i := range g.Grid {
// 		for j := range g.Grid[i] {
// 			if i && j {
// 				cells = append(cells, [2]{i,j})
// 			}
// 		}
// 	}
// }

// func checkIfAlive(grid [][]int, location [2]int) bool {
// 	n := countNeighbors(grid, location)
// 	if n == 2 || n == 3 {
// 		return true
// 	}
// 	return false
// }

func (g *Game) Iterate() error {
	return nil
}

func (g *Game) Print() {
	fmt.Print(g.Grid)
}

func (g *Game) Randomize() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			g.Grid[i][j] = r.Intn(2)
		}
	}
}

// returns a new Game struct with y rows and x cells per row
func NewGame(x int, y int) *Game {
	grid := make([][]int, y)
	row := make([]int, x)
	for i := range grid {
		grid[i] = row
	}
	return &Game{Grid: grid}
}
