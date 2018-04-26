package minimax

type Scorer struct {
	MaximizingPlayer Player
	MinimizingPlayer Player
}

func (scorer *Scorer) Score(game Game, player Player) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1
	} else if game.FindWinner() == scorer.MinimizingPlayer {
		return -1
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, move := range game.AvailableMoves() {
		nextGame := game.Move(move)
		nextScore := scorer.Score(nextGame, scorer.MinimizingPlayer)
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

type Game interface {
	IsOver() bool
	FindWinner() Player
	AvailableMoves() []Move
	Move(move Move) Game
}

type Move interface {}

type Player interface {}
