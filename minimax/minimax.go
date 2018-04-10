package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * player.Polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * player.Polarity
	} else if game.IsOver() {
		return 0
	}

	panic("no result")
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Player struct {
	Name string
	Polarity int
}
