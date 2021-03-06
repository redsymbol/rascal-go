package main

import (
	"code.google.com/p/goncurses"
)

type HandlerMapType map[goncurses.Key]func(*World)

func handler_move_left(world *World) {
	world.HandlePlayerMove(0, -1)
}

func handler_move_right(world *World) {
	world.HandlePlayerMove(0, 1)
}

func handler_move_up(world *World) {
	world.HandlePlayerMove(-1, 0)
}

func handler_move_down(world *World) {
	world.HandlePlayerMove(1, 0)
}

var HandlerMap HandlerMapType = HandlerMapType{
	'h': handler_move_left,
	'j': handler_move_down,
	'k': handler_move_up,
	'l': handler_move_right,
}
