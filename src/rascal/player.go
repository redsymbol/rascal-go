package main

type Player Actor

func NewPlayer(name string) * Player{
	return &Player{
		Name	: name,
		Symbol	: '@',
	}
}
