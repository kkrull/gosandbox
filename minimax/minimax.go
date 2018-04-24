package minimax

type Minimax struct {
	MaximizingPlayer Player
	MinimizingPlayer Player
}

func (scorer Minimax) Score(game Game, player Player) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1
	} else if game.FindWinner() == scorer.MinimizingPlayer {
		return -1
	} else if game.IsOver() {
		return 0
	}

	return 999
}

type Game interface {
	FindWinner() Player
	IsOver() bool
}

type Player struct {
	Name string
}
