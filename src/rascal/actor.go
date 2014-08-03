package main

import (
	"code.google.com/p/goncurses"
)

type Actor struct {
	Name   string
	Symbol goncurses.Char
	Point
}
