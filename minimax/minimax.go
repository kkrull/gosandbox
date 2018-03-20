package minimax

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

func (board Board) Minimax(player string) (string, int) {
	for _, space := range board.Spaces {
		if board.FindWinner(player, space) == player {
			if player == board.maxPlayer {
				return space.Id(), 1
			} else {
				return space.Id(), -1
			}
		}
	}

	for _, space := range board.Spaces {
		if space.IsAvailable() {
			return space.Id(), 0
		}
	}

	return "", 42
}

type Space interface {
	Id() string
	IsAvailable() bool
}
