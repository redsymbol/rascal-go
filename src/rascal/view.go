package main

import (
	"code.google.com/p/goncurses"
)

type View struct {
	Player *Actor
	World  *World
	Window *goncurses.Window
}

func NewView(player *Actor) *View {
	world := NewWorld(player)
	return &View{
		Player: player,
		World:  world,
		Window: nil,
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

	for {
		ch = window.GetChar()
		if ch == 'q' {
			break
		}

		if handler, ok := HandlerMap[ch]; ok {
			handler(view.World)
		}
		view.World.RunMonsterActions()
		if view.World.Player.Hitpoints <= 0 {
			return
		}
		view.Paint()
	}
}

func (view *View) InitPaint() {
	view.World.PositionActors()
	view.Paint()
}

func (view *View) Paint() {
	window := view.Window
	// draw terrain
	for xx, row := range view.World.Terrain {
		for yy, symbol := range row {
			window.Move(xx, yy)
			window.DelChar()
			window.AddChar(symbol)
		}
	}
	view.PaintDashboard()
	view.PaintMessage()
	view.PaintActors()
	window.Refresh()
}

func (view *View) PaintDashboard() {
	start_x := view.World.Height + 1
	start_y := 0
	view.ClearLine(start_x)
	view.Window.Move(start_x, start_y)
	view.Window.Printf("HP %d", view.World.Player.Hitpoints)
}

func (view *View) PaintMessage() {
	start_x := view.World.Height + 2
	start_y := 0
	view.Window.Move(start_x, start_y)
	view.ClearLine(start_x)
	view.Window.Printf(view.World.GetMessage())
}

func (view *View) ClearLine(lineno int) {
	view.Window.HLine(lineno, 0, ' ', view.World.Width)
}

func (view *View) PaintActors() {
	window := view.Window
	player := view.Player
	// draw monsters
	for _, monster := range view.World.AliveMonsters() {
		window.MoveAddChar(monster.X, monster.Y, monster.Symbol)
		window.Move(monster.X, monster.Y)
	}
	// draw player
	window.MoveAddChar(player.X, player.Y, player.Symbol)
	window.Move(player.X, player.Y)
}
