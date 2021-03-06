package main

import (
	"code.google.com/p/goncurses"
	"fmt"
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
	Player         *Actor
	Monsters       []*Actor
	Terrain        [][]goncurses.Char
	Width, Height  int
	Message        string
	currentMessage string
}

func NewWorld(player *Actor) *World {
	world := new(World)
	world.Player = player
	world.Monsters = make([]*Actor, 0, 0)
	world.TerrainFromMapping(_WORLD_MAPPING)
	world.currentMessage = ""
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

func (world *World) MoveActorBy(mover Mover, xdiff, ydiff int) {
	current_x, current_y := mover.GetPosition()
	new_x := current_x + xdiff
	new_y := current_y + ydiff
	if world.Occupiable(new_x, new_y) {
		mover.SetPosition(new_x, new_y)
	}
}

func (world *World) Occupiable(xx, yy int) bool {
	if world.Terrain[xx][yy] == '#' {
		return false
	}
	for _, monster := range world.AliveMonsters() {
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
	monsters := [...]*Actor{
		NewRat(),
		NewRat(),
		NewRat(),
		NewGiantRat(),
		NewGiantRat(),
		NewGoblin(),
		NewGoblin(),
		NewGoblin(),
	}
	for _, monster := range monsters {
		monster.X, monster.Y = world.findFreePosition()
		world.Monsters = append(world.Monsters, monster)
	}
}

func (world *World) AliveMonsters() []*Actor {
	monsters := make([]*Actor, 0, 0)
	for _, monster := range world.Monsters {
		if monster.IsAlive() {
			monsters = append(monsters, monster)
		}
	}
	return monsters
}

func (world *World) RunMonsterActions() {
	for _, monster := range world.AliveMonsters() {
		if monster.AdjacentTo(world.Player.X, world.Player.Y) {
			world.AddMessage(fmt.Sprintf("%s hits you for %d damage!", monster.Name, monster.Damage))
			monster.Attack(world.Player)
		} else {
			var dx, dy, abs_dx, abs_dy, offset_x, offset_y int
			dx = world.Player.X - monster.X
			dy = world.Player.Y - monster.Y
			if dx < 0 {
				abs_dx = -1 * dx
			} else {
				abs_dx = dx
			}
			if dy < 0 {
				abs_dy = -1 * dy
			} else {
				abs_dy = dy
			}
			if abs_dx > abs_dy {
				offset_y = 0
				if dx > 0 {
					offset_x = 1
				} else {
					offset_x = -1
				}
			} else {
				offset_x = 0
				if dy > 0 {
					offset_y = 1
				} else {
					offset_y = -1
				}
			}
			world.MoveActorBy(monster, offset_x, offset_y)
		}

	}
}

func (world *World) AddMessage(msg string) {
	world.currentMessage = msg
}

func (world *World) GetMessage() string {
	return world.currentMessage
}

func (world *World) HandlePlayerMove(offset_x, offset_y int) {
	target_x := world.Player.X + offset_x
	target_y := world.Player.Y + offset_y
	var monster *Actor = nil
	for _, this_monster := range world.AliveMonsters() {
		if this_monster.X == target_x && this_monster.Y == target_y {
			monster = this_monster
			break
		}
	}
	if monster == nil {
		world.MoveActorBy(world.Player, offset_x, offset_y)
	} else {
		world.AddMessage(fmt.Sprintf("You hit the %s for %d damage.", monster.Name, world.Player.Damage))
		world.Player.Attack(monster)
		if !monster.IsAlive() {
			world.AddMessage(fmt.Sprintf("You have slain the %s!", monster.Name))
		}
	}
}

func unit_of(num int) int {
	if num > 0 {
		num = 1
	} else if num < 0 {
		num = -1
	}
	return num
}

func abs(num int) int {
	if num < 1 {
		return -1 * num
	}
	return num
}
