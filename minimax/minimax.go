package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else {
		return -1
	}
}

type Game interface {
	FindWinner() Player
	MaximizingPlayer() Player
}

type Player struct {
	Name string
}
