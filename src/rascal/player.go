package main

import (
	"code.google.com/p/goncurses"
)

type Player struct {
	Name string
	X, Y int
}

func (p *Player) Symbol() goncurses.Char {
	return '@'
}
