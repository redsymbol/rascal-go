package main

type Monster Actor

func NewRat() *Monster {
	return &Monster{
		Name:      "rat",
		Symbol:    'r',
	}
}

func NewGiantRat() *Monster {
	return &Monster{
		Name:      "giant rat",
		Symbol:    'R',
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
