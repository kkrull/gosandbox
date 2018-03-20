package minimax

type Board struct {
	Spaces []Space
}

func (board Board) Minimax() string {
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
