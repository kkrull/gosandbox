package minimax

import "errors"

func Minimax(game Game) (result Move, err error) {
	return "", errors.New("minimax: game over")
}

type Move interface {
}

type Game interface {
}