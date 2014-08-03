package main

import (
	"code.google.com/p/goncurses"
	"math/rand"
	"strings"
)

var _WORLD_MAPPING string = `
############################################################
#                                                          #
#                                                          #
#                                                          #
#                                                          #
#              ##########                                  #
#              #        #                                  #
#                       ###########################        #
#              #                                  #        #
#              ########################           #        #
#                                     #           #        #
#                                     #############        #
#############                                              #
#                                                          #
#           #                                              #
#           #                                              #
#           #                                              #
#############                                              #
#                                                          #
#                                                          #
############################################################
`

type World struct {
	Player        *Player
	Monsters      []*Monster
	Terrain       [][]goncurses.Char
	Width, Height int
}

func NewWorld(player *Player) *World {
	world := new(World)
	world.Player = player
	world.Monsters = make([]*Monster, 0, 0)
	world.TerrainFromMapping(_WORLD_MAPPING)
	return world
}

func (world *World) TerrainFromMapping(mapping string) {
	trimmed_mapping := strings.Trim(mapping, "\n\t ")
	world.Height = strings.Count(trimmed_mapping, "\n") + 1
	world.Width = -1
	world.Terrain = make([][]goncurses.Char, world.Height, world.Height)
	var xx, yy int
	var line string
	var tile rune
	for xx, line = range strings.Split(trimmed_mapping, "\n") {
		if world.Width == -1 {
			world.Width = len(line)
		} else if len(line) != world.Width {
			panic("Inconsistent line lengths in mapping")
		}
		world.Terrain[xx] = make([]goncurses.Char, world.Width, world.Width)
		for yy, tile = range line {
			world.Terrain[xx][yy] = goncurses.Char(tile)
		}
	}
}

func (world *World) MovePlayerTo(xdiff, ydiff int) {
	new_x := world.Player.X + xdiff
	new_y := world.Player.Y + ydiff
	if world.Occupiable(new_x, new_y) {
		world.Player.X = new_x
		world.Player.Y = new_y
	}
}

func (world *World) Occupiable(xx, yy int) bool {
	if world.Terrain[xx][yy] == '#' {
		return false
	}
	for _, monster := range world.Monsters {
		if monster.X == xx && monster.Y == yy {
			return false
		}
	}
	return true
}

func (world *World) findFreePosition() (xx, yy int) {
	for {
		xx = rand.Intn(world.Height)
		yy = rand.Intn(world.Width)
		if world.Occupiable(xx, yy) {
			break
		}
	}
	return
}

func (world *World) PositionActors() {
	world.Player.X, world.Player.Y = world.findFreePosition()
	monsters := [...]*Monster{
		NewRat(),
		NewGiantRat(),
	}
	for _, monster := range monsters {
		monster.X, monster.Y = world.findFreePosition()
		world.Monsters = append(world.Monsters, monster)
	}
}
