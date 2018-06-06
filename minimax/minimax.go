package minimax

type Scorer struct {
	Maximizer Player
	Minimizer Player
}

func (scorer *Scorer) Score(game GameState, player Player) int {
	if player == scorer.Maximizer {
		return scorer.Negamax(game, 1)
	} else {
		return -scorer.Negamax(game, -1)
	}
}

func (scorer *Scorer) Negamax(game GameState, polarity int) int {
	if game.FindWinner() == scorer.Maximizer {
		return 1 * polarity
	} else if game.FindWinner() == scorer.Minimizer {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, nextMove := range game.AvailableMoves() {
		nextGame := game.Move(nextMove)
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
