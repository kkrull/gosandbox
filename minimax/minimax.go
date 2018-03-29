package minimax

func Minimax(game Game, player string) Result {
	switch game.FindWinner() {
	case "Max":
		return Result{Score: 1}
	case "Min":
		return Result{Score: -1}
	default:
		return Result{Score: 0}
	}
}

type Game interface {
	FindWinner() string
}

type Result struct {
	Score int
}
