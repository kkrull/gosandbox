package minimax

type Scorer struct {
}

func (scorer *Scorer) Score(game GameState, player Player) int {
	panic("Not implemented")
}

type GameState interface {
}

type Move string
type Player string
