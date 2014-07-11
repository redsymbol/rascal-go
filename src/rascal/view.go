package main
import "fmt"
type View struct {
	Player Player;
}

func (view *View) RunForever() {
	fmt.Printf("Hello, %s.\n", view.Player.Name)
}
