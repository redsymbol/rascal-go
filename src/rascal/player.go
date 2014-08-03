package main

type Player ActorBase

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Symbol: '@',
		Hitpoints: 2,
		Damage: 1,
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

func (player *Player) GetDamage() int {
	return player.Damage 
}

func (player *Player) Attack(other *Monster) {
	other.ReduceHitpoints(player.GetDamage())
}

func (player *Player) ReduceHitpoints(damage int) {
	player.Hitpoints -= damage
}
