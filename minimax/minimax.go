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
	if space, score := board.findWinningSpace(player); space != nil {
		return moveToSpace(space, score)
	}

	for _, space := range board.Spaces {
		if space.IsAvailable() {
			return moveToSpace(space, 0)
		}
	}

	return gameOver()
}

func (board Board) findWinningSpace(player string) (Space, int) {
	for _, space := range board.Spaces {
		if board.FindWinner(player, space) == player {
			if player == board.maxPlayer {
				return space, 1
			} else {
				return space, -1
			}
		}
	}

	return nil, 0
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
