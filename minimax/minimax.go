package minimax

func Minimax(game Game, player string) Result {
	switch game.FindWinner() {
	case "Max":
		return Result{Score: 1}
	case "Min":
		return Result{Score: -1}
	}

	if game.IsOver() {
		return Result{Score: 0}
	}

	panic("No decision")
}

type Game interface {
	FindWinner() string
	IsOver() bool
}

type Result struct {
	Score int
	Space string
}
