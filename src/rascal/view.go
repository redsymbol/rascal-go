package main
import (
	"code.google.com/p/goncurses"
)

type View struct {
	Player *Player
	World *World
	Window *goncurses.Window
}

func NewView(player *Player) *View {
	world := NewWorld(player)
	return &View{
		player,
		world,
		nil,
	}
}

func (view *View) RunForever() {
	window, err := goncurses.Init()
	view.Window = window
	if err != nil {
		panic(err)
	}
	defer goncurses.End()
	
	goncurses.CBreak(true)
	goncurses.Echo(false)
	window.Clear()
	var ch goncurses.Key

	view.InitPaint()

	for ;; {
		ch = window.GetChar()
		if ch == 'q' {
			break
		}
		view.Paint()
		
	}
}

func (view *View) InitPaint() {
	for xx, row := range view.World.Terrain {
		for yy, symbol := range row {
			view.Window.Move(xx, yy)
			view.Window.DelChar()
			view.Window.AddChar(symbol)
		}
	}
}

func (view *View) Paint() {
}
