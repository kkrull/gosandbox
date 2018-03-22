package minimax

import (
	"errors"
)

func Minimax(player Player, game Game) (result Move, err error) {
	if game.IsOver() {
		return nil, errors.New("minimax: game over")
	}

	for _, space := range game.OpenSpaces() {
		gameWouldBe := game.ClaimSpace(player, space)
		if gameWouldBe.FindWinner() == player {
			return space, nil
		}
	}

	return nil, nil
}

type Move interface {}

type Game interface {
	ClaimSpace(Player, Space) Game
	FindWinner() Player
	IsOver() bool
	OpenSpaces() []Space
}

type Player interface {}

type Space interface{}
