package main
import "fmt"
type View struct {
	Player *Player;
	World *World;
}

func (view *View) RunForever() {
	fmt.Printf("Hello, %s.\n", view.Player.Name)
}

func NewView(player *Player) *View {
	world := NewWorld(player)
	return &View{
		player,
		world,
	}
}
