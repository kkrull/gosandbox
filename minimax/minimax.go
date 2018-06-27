package minimax

type Scorer struct {
	Maximizer Player
	Minimizer Player
}

func (scorer *Scorer) Score(game GameState, player Player) int {
	if player == scorer.Maximizer {
		return scorer.negamax(game, 1)
	} else {
		return reversePerspective(scorer.negamax(game, -1))
	}
}

func (scorer *Scorer) negamax(game GameState, polarity int) int {
	if game.FindWinner() == scorer.Maximizer {
		return 1 * polarity
	} else if game.FindWinner() == scorer.Minimizer {
		return reverseThe(polarity)
	} else if game.IsOver() {
		return 0
	}

	maxScore := -10
	for _, nextMove := range game.AvailableMoves() {
		nextGame := game.Move(nextMove)
		nextScore := reversePerspective(scorer.negamax(nextGame, reverseThe(polarity)))
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

func reversePerspective(score int) int {
	return -1 * score
}

func reverseThe(polarity int) int {
	return -1 * polarity
}

type GameState interface {
	AvailableMoves() []Move
	FindWinner() Player
	IsOver() bool
	Move(move Move) GameState
}

type Move string
type Player string
