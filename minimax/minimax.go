package minimax

type Scorer struct {
	MaximizingPlayer Player
	MinimizingPlayer Player
}

func (scorer *Scorer) Score(game GameState, player Player) int {
	if player == scorer.MaximizingPlayer {
		return scorer.Negamax(game, 1)
	} else {
		return -scorer.Negamax(game, -1)
	}
}

func (scorer *Scorer) Negamax(game GameState, polarity int) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1 * polarity
	} else if game.FindWinner() == scorer.MinimizingPlayer {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, move := range game.AvailableMoves() {
		nextGame := game.Move(move)
		nextScore := -scorer.Negamax(nextGame, -1*polarity)
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

type GameState interface {
	AvailableMoves() []Move
	IsOver() bool
	FindWinner() Player
	Move(move Move) GameState
}

type Move string
type Player string
