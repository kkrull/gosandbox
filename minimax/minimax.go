package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.IsOver() {
		return 0
	}

	return 999
}

type Game interface {
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
}

type Player struct {
	Name string
}
