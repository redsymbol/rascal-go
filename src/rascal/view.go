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
	window := view.Window
	player := view.Player
	// draw terrain
	for xx, row := range view.World.Terrain {
		for yy, symbol := range row {
			window.Move(xx, yy)
			window.DelChar()
			window.AddChar(symbol)
		}
	}
	// position actors
	player.X = 5
	player.Y = 5
	// draw player
	window.MoveAddChar(player.X, player.Y, player.Symbol())
	window.Move(player.X, player.Y)
	window.Refresh()
}

func (view *View) Paint() {
}
