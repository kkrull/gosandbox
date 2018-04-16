package minimax

func Minimax(game Game, player Player) int {
	if player == game.MaximizingPlayer() {
		return Negamax(game, 1)
	} else {
		return -Negamax(game, -1)
	}
}

func Negamax(game Game, polarity int) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	bestScore := -100
	for _, nextGame := range game.NextGames() {
		scoreFromOpponentPerspective := Negamax(nextGame, -1 * polarity)
		scoreFromMyPerspective := -1 * scoreFromOpponentPerspective
		if scoreFromMyPerspective > bestScore {
			bestScore = scoreFromMyPerspective
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
