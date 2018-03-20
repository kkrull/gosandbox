package minimax

import "fmt"

//import "fmt"

func NewBoard(maxPlayer string, minPlayer string, spaces []Space, findWinner func(string, Space) string) Board {
	return Board{
		maxPlayer:  maxPlayer,
		minPlayer:  minPlayer,
		FindWinner: findWinner,
		Spaces:     spaces,
	}
}

type Board struct {
	maxPlayer  string
	minPlayer  string
	FindWinner func(ifPlayer string, claimsSpace Space) string
	Spaces     []Space
}

func (board Board) Minimax(player string) (string, int, error) {
	for _, space := range board.Spaces {
		if board.FindWinner(player, space) == player {
			if player == board.maxPlayer {
				return moveToSpace(space, 1)
			} else {
				return moveToSpace(space, -1)
			}
		}
	}

	for _, space := range board.Spaces {
		if space.IsAvailable() {
			return moveToSpace(space, 0)
		}
	}

	return gameOver()
}

func moveToSpace(space Space, score int) (string, int, error) {
	return space.Id(), score, nil
}

func gameOver() (string, int, error) {
	return "", 0, fmt.Errorf("the game is already over")
}

type Space interface {
	Id() string
	IsAvailable() bool
}
