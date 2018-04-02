package minimax

func Minimax(game Game) Result {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Result{Score: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Result{Score: -1}
	}
	return Result{}
}

type Game interface {
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Player struct {
	Name string
}

type Result struct {
	Score int
}
