package minimax

type Board struct {
	Spaces []Space
}

func (board Board) Minimax() {
	
}

type Space interface {
	Id() string
	//IsAvailable() bool
}
