package minimax

func Minimax(game Game, player Player) int {
	return Negamax(game, 1)
}

func Negamax(game Game, polarity int) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return polarity * 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return polarity * -1
	} else if game.IsOver() {
		return 0
	}

	bestScore := -100
	for _, nextGame := range game.NextGames() {
		nextScoreFromOpponentPerspective := Negamax(nextGame, -1 * polarity)
		nextScoreFromMyPerspective := -nextScoreFromOpponentPerspective
		if nextScoreFromMyPerspective > bestScore {
			bestScore = nextScoreFromMyPerspective
		}
	}

	return bestScore
}

type Game interface {
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	NextGames() []Game
}

type Player struct {
	Name string
}
