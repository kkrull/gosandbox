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

	move := game.LegalMoves(player)[0]
	nextResult := Minimax(move.Game, player)
	return Result{Score: nextResult.Score, Space: move.ClaimedSpace}
}

type Game interface {
	FindWinner() string
	IsOver() bool
	LegalMoves(player string) []Move
}

type Move struct {
	Game Game
	ClaimedSpace string
}

type Result struct {
	Score int
	Space string
}
