package main

import (
	"code.google.com/p/goncurses"
)

type Mover interface {
	SetPosition(x, y int)
	GetPosition() (x, y int)
	AdjacentTo(x, y int) bool
}

type Combatter interface {
	Attack(other *Combatter)
	ReduceHitpoints(damage int)
	GetDamage() int
}

type ActorBase struct {
	Name      string
	Symbol    goncurses.Char
	Hitpoints int
	Damage    int
	Point
}

