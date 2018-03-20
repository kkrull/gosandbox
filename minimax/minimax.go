package minimax

//import "fmt"

func NewBoard(spaces []Space, findWinner func(string, Space) string) Board {
	return Board{Spaces: spaces, FindWinner: findWinner}
}

type Board struct {
	Spaces     []Space
	FindWinner func(ifPlayer string, claimsSpace Space) string
}

func (board Board) Minimax(player string) string {
	for _, space := range board.Spaces {
		if board.FindWinner(player, space) == player {
			return space.Id()
		}
	}

	for _, space := range board.Spaces {
		if space.IsAvailable() {
			return space.Id()
		}
	}

	return ""
}

type Space interface {
	Id() string
	IsAvailable() bool
}
