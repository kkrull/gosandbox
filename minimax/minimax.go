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

	maxScore := -100
	for _, nextGame := range game.NextPossibleGames() {
		nextScore := scorer.Score(nextGame, scorer.MinimizingPlayer)
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	NextPossibleGames() []Game
}

type Player struct {
	Name string
}
