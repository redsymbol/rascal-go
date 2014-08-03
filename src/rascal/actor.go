package main

import (
	"code.google.com/p/goncurses"
)

type Mover interface {
	SetPosition(x, y int)
	GetPosition() (x, y int)
	AdjacentTo(x, y int) bool
}

type Actor struct {
	Name   string
	Symbol goncurses.Char
	Point
}
