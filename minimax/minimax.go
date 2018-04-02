package minimax

func Minimax(game Game) Result {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Result{Score: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Result{Score: -1}
	} else if game.IsOver() {
		return Result{Score: 0}
	}

	bestResult := Result{Score: -100}
	for _, outcome := range game.AvailableMoves() {
		nextResult := Minimax(outcome.NextGame)
		if nextResult.Score > bestResult.Score {
			bestResult = Result{
				Score: nextResult.Score,
				Move:  outcome.Move}
		}
	}

	return bestResult
}

type Game interface {
	AvailableMoves() []Outcome
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Outcome struct {
	Move     Move
	NextGame Game
}

type Move interface {
}

type Player struct {
	Name string
}

type Result struct {
	Move  Move
	Score int
}
