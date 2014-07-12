package main

import (
	"strings"
	"code.google.com/p/goncurses"
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
	Player *Player
	Terrain [][]goncurses.Char
}

func NewWorld(player *Player) *World{
	world := new(World)
	world.Player = player
	world.TerrainFromMapping(_WORLD_MAPPING)
	return world
}

func (world *World) TerrainFromMapping(mapping string) {
	var height int
	var width int
	trimmed_mapping := strings.Trim(mapping, "\n\t ")
	height = strings.Count(trimmed_mapping, "\n") + 1
	width = -1
	world.Terrain = make([][]goncurses.Char, height, height)
	for xx, line := range strings.Split(trimmed_mapping, "\n") {
		if width == -1 {
			width = len(line)
		} else if len(line) != width {
			panic("Inconsistent line lengths in mapping")
		}
		world.Terrain[xx] = make([]goncurses.Char, width, width)
		for yy, tile := range line {
			world.Terrain[xx][yy] = goncurses.Char(tile)
		}
	}
}
