package main

type Monster Actor

func NewRat() *Monster {
	return &Monster{
		Name:   "rat",
		Symbol: 'r',
	}
}

func NewGiantRat() *Monster {
	return &Monster{
		Name:   "giant rat",
		Symbol: 'R',
	}
}
