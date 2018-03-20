package minimax

//import "fmt"

func NewBoard(spaces []Space, findWinner func(Space) string) Board {
	return Board{Spaces: spaces, FindWinner: findWinner}
}

type Board struct {
	Spaces     []Space
	FindWinner func(spaceToTake Space) string
}

func (board Board) Minimax() string {
	for _, space := range board.Spaces {
		if board.FindWinner(space) == "max" {
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
