package minimax

import "fmt"

func Minimax(game Game) (result string, err error) {
	if game.IsOver() {
		return "", fmt.Errorf("the game is already over")
	}

	return "Victory", nil
	//return 0, fmt.Errorf("unable to pick a space")
}

type Game interface {
	IsOver() bool
}
