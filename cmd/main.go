package main

import (
	"github.com/aglovaile/golang-life/pkg/life"
)

func main() {
	l := life.NewGame(10, 10)
	l.Print()
}
