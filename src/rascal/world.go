package main

import (
	"code.google.com/p/goncurses"
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
	Player  *Player
	Terrain [][]goncurses.Char
	Width, Height int
}

func NewWorld(player *Player) *World {
	world := new(World)
	world.Player = player
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
