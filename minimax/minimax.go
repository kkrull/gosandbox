package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, successor := range game.Successors() {
		score := Negamax(successor, player)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	Successors() []Game
}

type Player struct {
	Name string
}
