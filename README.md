[![CircleCI](https://dl.circleci.com/status-badge/img/gh/aglovaile/golang-life/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/aglovaile/golang-life/tree/master)

# golang-life
A Golang implementation of Conway's Game of Life

## `life` package:

Exports:
- `Game` struct
- `NewGame(x int, y int)` func - returns a new Game with 0 for all Game.Grid cells and for Game.Iterations 
- `NewGameFromSeed(seed [][]int)` func - returns a new Game with Game.Grid equal to provided seed. Seed must be a 2D array of rows which are __equal in length__ and __have values of 0 or 1__.

It is recommended to use NewGame() or NewGameFromSeed() to instantialize a new Game struct rather than calling &Game directly.

## `Game` struct:

Fields:
- `Game.Grid` - contains a 2D array of the game with __each cell having 1 for alive and 0 for dead__
- `Game.Iterations` - a counter for how many times the Game has been iterated

Methods:
- `Game.Iterate` - advances the Game.Grid forward one generation
- `Game.Randomize` - fills the Game.Grid with random values