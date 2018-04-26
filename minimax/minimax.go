package minimax

type Scorer struct {

}

func (scorer *Scorer) Score(game Game, player Player) int {
	return 999
}

type Game interface {

}

type Player interface {}
