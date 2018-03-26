package minimax

import "fmt"

func Minimax(game Game) (result int, err error) {
	if game.IsOver() {
		return 0, fmt.Errorf("the game is already over")
	}

	return 0, fmt.Errorf("unable to pick a space")
}

type Game interface {
	IsOver() bool
}
