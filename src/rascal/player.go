package main

import (
	"code.google.com/p/goncurses"
)

type Player struct {
	Name string
	Point
}

func (p *Player) Symbol() goncurses.Char {
	return '@'
}
