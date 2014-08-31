package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	view := NewView(NewPlayer("Aaron"))
	exit_message := view.RunForever()
	if exit_message != "" {
		fmt.Println(exit_message)
	}
}
