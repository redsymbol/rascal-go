package main

type Player Actor

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Symbol: '@',
	}
}

func (player *Player) SetPosition(x, y int) {
	player.X = x
	player.Y = y
}

func (player *Player) GetPosition() (x, y int) {
	x = player.X
	y = player.Y
	return
}

func (player *Player) AdjacentTo(x, y int) bool {
	return PointsAdjacent(player.X, player.Y, x, y)
}
