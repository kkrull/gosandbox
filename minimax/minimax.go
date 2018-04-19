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

	maxScore := -100
	for _, nextGame := range game.NextGames() {
		nextScore := -Negamax(nextGame, -1 * polarity)
		if nextScore > maxScore {
			maxScore = nextScore
		}
	}

	return maxScore
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	NextGames() []Game
}

type Player struct {
	Name string
}
