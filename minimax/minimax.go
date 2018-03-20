package minimax

type Board struct {
	Spaces []Space
}

type Space interface {
	Id() string
	//IsAvailable() bool
}
