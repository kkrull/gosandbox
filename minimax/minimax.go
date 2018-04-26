package minimax

type Scorer struct {
	MaximizingPlayer Player
}

func (scorer *Scorer) Score(game Game, player Player) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1
	} else if game.IsOver() {
		return 0
	}

	return 999
}

type Game interface {
	IsOver() bool
	FindWinner() Player
}

type Player interface {}
