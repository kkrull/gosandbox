package minimax

type Board struct {
	Spaces []Space
}

func (board Board) Minimax() string {
	return "A1"
}

type Space interface {
	Id() string
	//IsAvailable() bool
}
