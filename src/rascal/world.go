package main

import (
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
	Player *Player
	terrain [][]rune
}

func NewWorld(player *Player) *World{
	world := new(World)
	world.Player = player
	world.terrainFromMapping(_WORLD_MAPPING)
	return world
}

func (world *World) terrainFromMapping(mapping string) {
	var height int
	var width int
	trimmed_mapping := strings.Trim(mapping, "\n\t ")
	height = strings.Count(trimmed_mapping, "\n") + 1
	width = -1
	world.terrain = make([][]rune, height, height)
	for xx, line := range strings.Split(trimmed_mapping, "\n") {
		if width == -1 {
			width = len(line)
		} else if len(line) != width {
			panic("Inconsistent line lengths in mapping")
		}
		world.terrain[xx] = make([]rune, width, width)
		for yy, tile := range line {
			world.terrain[xx][yy] = tile
		}
	}
}
