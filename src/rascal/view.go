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
	window := view.Window
	// draw terrain
	for xx, row := range view.World.Terrain {
		for yy, symbol := range row {
			window.Move(xx, yy)
			window.DelChar()
			window.AddChar(symbol)
		}
	}
	// position and initially paint actors
	view.World.PositionActors()
	view.PaintActors()
	window.Refresh()
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
	view.PaintActors()
	window.Refresh()
}

func (view *View) PaintActors() {
	window := view.Window
	player := view.Player
	// draw monsters
	for _, monster := range view.World.Monsters {
		window.MoveAddChar(monster.X, monster.Y, monster.Symbol)
		window.Move(monster.X, monster.Y)
	}
	// draw player
	window.MoveAddChar(player.X, player.Y, player.Symbol)
	window.Move(player.X, player.Y)
}
