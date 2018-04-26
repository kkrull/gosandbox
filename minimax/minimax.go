package minimax

type Scorer struct {

}

func (scorer *Scorer) Score(game Game, player Player) int {
	if game.IsOver() {
		return 0
	}

	return 999
}

type Game interface {
	IsOver() bool
}

type Player interface {}
