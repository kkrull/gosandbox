package minimax

type Scorer struct {
	MaximizingPlayer Player
	MinimizingPlayer Player
}

func (scorer Scorer) Score(game GameState, player Player) int {
	if player == scorer.MaximizingPlayer {
		return scorer.Negamax(game, 1)
	} else {
		return -scorer.Negamax(game, -1)
	}
}

func (scorer Scorer) Negamax(game GameState, polarity int) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1 * polarity
	} else if game.FindWinner() == scorer.MinimizingPlayer {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, move := range game.AvailableMoves() {
		nextGameState := game.NextGameState(move)
		nextScore := -scorer.Negamax(nextGameState, -1*polarity)
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

type GameState interface {
	AvailableMoves() []Move
	FindWinner() Player
	IsOver() bool
	NextGameState(Move) GameState
}

type Move interface{}

type Player string
