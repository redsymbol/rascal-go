package main

func main() {
	view := NewView(&Player{Name:"Aaron"})
	view.RunForever()
}
