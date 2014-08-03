package main

type Monster ActorBase

func NewRat() *Monster {
	return &Monster{
		Name:      "rat",
		Symbol:    'r',
		Hitpoints: 1,
		Damage:    1,
	}
}

func NewGiantRat() *Monster {
	return &Monster{
		Name:      "giant rat",
		Symbol:    'R',
		Hitpoints: 2,
		Damage:    1,
	}
}

func (monster *Monster) SetPosition(x, y int) {
	monster.X = x
	monster.Y = y
}

func (monster *Monster) GetPosition() (x, y int) {
	x = monster.X
	y = monster.Y
	return
}

func (monster *Monster) AdjacentTo(x, y int) bool {
	return PointsAdjacent(monster.X, monster.Y, x, y)
}

func (monster *Monster) GetDamage() int {
	return monster.Damage 
}

func (monster *Monster) Attack(other *Player) {
	other.ReduceHitpoints(monster.GetDamage())
}

func (monster *Monster) ReduceHitpoints(damage int) {
	monster.Hitpoints -= damage
}
