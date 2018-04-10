package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	}

	return 0
}

type Game interface {
	FindWinner() Player
	MaximizingPlayer() Player
}

type Player struct {
	Name string
}
