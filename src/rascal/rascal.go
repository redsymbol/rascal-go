package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	view := NewView(NewPlayer("Aaron"))
	view.RunForever()
}
