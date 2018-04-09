package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return coloredByPlayer(1, player, game)
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return coloredByPlayer(-1, player, game)
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	opponent := opponentOf(player, game)
	for _, successor := range game.Successors() {
		score := -Negamax(successor, opponent)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func coloredByPlayer(score int, player Player, game Game) int {
	if player == game.MaximizingPlayer() {
		return score
	} else {
		return -score
	}
}

func opponentOf(player Player, game Game) Player {
	if player == game.MaximizingPlayer() {
		return game.MinimizingPlayer()
	} else {
		return game.MaximizingPlayer()
	}
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
