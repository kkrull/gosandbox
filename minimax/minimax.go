package minimax

func Minimax(game Game) Result {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Result{Score: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Result{Score: -1}
	} else if game.IsOver() {
		return Result{Score: 0}
	}

	panic("no result")
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Move interface {
}

type Player struct {
	Name string
}

type Result struct {
	Move Move
	Score int
}
