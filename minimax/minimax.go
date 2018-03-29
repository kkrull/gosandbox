package minimax

func Minimax(game Game, player string) Result {
	return Result{Score: 1}
}

type Game interface {

}

type Result struct {
	Score int
}
