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
func (monster *Actor) SetPosition(x, y int) {
	monster.X = x
	monster.Y = y
}

func (monster *Actor) GetPosition() (x, y int) {
	x = monster.X
	y = monster.Y
	return
}

func (monster *Actor) AdjacentTo(x, y int) bool {
	return PointsAdjacent(monster.X, monster.Y, x, y)
}

func (monster *Actor) GetDamage() int {
	return monster.Damage 
}

func (monster *Actor) Attack(other *Actor) {
	other.ReduceHitpoints(monster.GetDamage())
}

func (monster *Actor) ReduceHitpoints(damage int) {
	monster.Hitpoints -= damage
}

// New actor constructors
func NewPlayer(name string) *Actor {
	return &Actor{
		Name:   name,
		Symbol: '@',
		Hitpoints: 10,
		Damage: 1,
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

