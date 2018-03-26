package minimax

import "fmt"

func Minimax(game Game) (space Space, err error) {
	if game.IsOver() {
		return "", fmt.Errorf("the game is already over")
	}

	space = game.OpenSpaces()[0]
	return space, nil
	//return 0, fmt.Errorf("unable to pick a space")
}

type Game interface {
	OpenSpaces() []Space
	IsOver() bool
}

type Space interface {
}
