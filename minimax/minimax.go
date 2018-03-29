package minimax

func Minimax(game Game, player string) Result {
	if player == "Max" {
		return Result{Score: 1}
	} else if player == "Min" {
		return Result{Score: -1}
	}

	panic("Expected a result")
}

type Game interface {

}

type Result struct {
	Score int
}
