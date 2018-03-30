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

	var bestMove Move
	var bestScore = -100
	for _, move := range game.LegalMoves(player) {
		nextResult := Minimax(move.Game, "Min")
		if nextResult.Score > bestScore {
			bestMove = move
			bestScore = nextResult.Score
		}
	}

	return Result{Score: bestScore, Space: bestMove.ClaimedSpace}
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
