package main
import (
	"code.google.com/p/goncurses"
)	

type HandlerMapType map[goncurses.Key]func(*World)

func handler_move_left(world *World) {
	world.Player.Y -= 1
}

func handler_move_right(world *World) {
	world.Player.Y += 1
}

func handler_move_up(world *World) {
	world.Player.X -= 1
}

func handler_move_down(world *World) {
	world.Player.X += 1
}

var HandlerMap HandlerMapType = HandlerMapType{
	'h' : handler_move_left,
	'j' : handler_move_down,
	'k' : handler_move_up,
	'l' : handler_move_right,
}
