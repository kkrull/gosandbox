package minimax

import "errors"

func Minimax(game Game) (result Move, err error) {
	if game.IsOver() {
		return nil, errors.New("minimax: game over")
	}

	return game.OpenSpaces()[0], nil
}

type Move interface {}

type Game interface {
	IsOver() bool
	OpenSpaces() []Space
}

type Space interface{}
