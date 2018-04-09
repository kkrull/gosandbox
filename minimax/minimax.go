package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else {
		return 0
	}
}

type Game interface {
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Player struct {
	Name string
}
