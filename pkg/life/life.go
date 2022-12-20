package life

import "fmt"

type Game struct {
	Grid       [][]int
	Iterations int
}

func (g *Game) Iterate() error {
	return nil
}

func (g *Game) Print() {
	fmt.Print(g.Grid)
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
