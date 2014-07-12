package main
import (
	"code.google.com/p/goncurses"
)

type View struct {
	Player *Player
	World *World
	window *goncurses.Window
}

func (view *View) RunForever() {
	window, err := goncurses.Init()
	if err != nil {
		panic(err)
	}
	defer goncurses.End()
	
	goncurses.CBreak(true)
	goncurses.Echo(false)
	window.Clear()
	var ch goncurses.Key

	for ;; {
		ch = window.GetChar()
		if ch == 'q' {
			break
		}
	}
}

func NewView(player *Player) *View {
	world := NewWorld(player)
	return &View{
		player,
		world,
		nil,
	}
}
