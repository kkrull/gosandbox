package minimax

func Minimax(game Game, player Player) int {
	return Negamax(game, 1)
}

func Negamax(game Game, polarity int) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	nextGame := game.NextGames()[0]
	return -Negamax(nextGame, -1 * polarity)
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
