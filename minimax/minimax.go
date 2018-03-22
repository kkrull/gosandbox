package minimax

import "errors"

func Minimax(game Game) (result Move, err error) {
	if game.IsOver() {
		return "", errors.New("minimax: game over")
	}

	return "", nil
}

type Move interface {
}

type Game interface {
	IsOver() bool
}
