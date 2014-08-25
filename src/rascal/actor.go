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

type Actor struct {
	Name      string
	Symbol    goncurses.Char
	Hitpoints int
	Damage    int
	Point
}

// Actor methods
func (actor *Actor) SetPosition(x, y int) {
	actor.X = x
	actor.Y = y
}

func (actor *Actor) GetPosition() (x, y int) {
	x = actor.X
	y = actor.Y
	return
}

func (actor *Actor) AdjacentTo(x, y int) bool {
	return PointsAdjacent(actor.X, actor.Y, x, y)
}

func (actor *Actor) GetDamage() int {
	return actor.Damage
}

func (actor *Actor) Attack(other *Actor) {
	other.ReduceHitpoints(actor.GetDamage())
}

func (actor *Actor) ReduceHitpoints(damage int) {
	actor.Hitpoints -= damage
}

func (actor *Actor) IsAlive() bool {
	return actor.Hitpoints > 0
}

// New actor constructors
func NewPlayer(name string) *Actor {
	return &Actor{
		Name:      name,
		Symbol:    '@',
		Hitpoints: 10,
		Damage:    1,
	}
}

func NewRat() *Actor {
	return &Actor{
		Name:      "rat",
		Symbol:    'r',
		Hitpoints: 1,
		Damage:    1,
	}
}

func NewGiantRat() *Actor {
	return &Actor{
		Name:      "giant rat",
		Symbol:    'R',
		Hitpoints: 2,
		Damage:    1,
	}
}

func NewGoblin() *Actor {
	return &Actor{
		Name:      "goblin",
		Symbol:    'g',
		Hitpoints: 1,
		Damage:    2,
	}
}
