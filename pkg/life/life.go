package life

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (g *Game) Randomize() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			if r.Intn(2) != 0 {
				g.Grid[i][j] = 1
			} else {
				g.Grid[i][j] = 0
			}
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
