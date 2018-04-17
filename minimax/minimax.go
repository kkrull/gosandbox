package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else if game.IsOver() {
		return 0
	}

	if player == game.MaximizingPlayer() {
		bestScore := -100
		for _, nextGame := range game.NextGames() {
			score := Minimax(nextGame, player)
			if score > bestScore {
				bestScore = score
			}
		}

		return bestScore
	} else {
		bestScore := 100
		for _, nextGame := range game.NextGames() {
			score := Minimax(nextGame, player)
			if score < bestScore {
				bestScore = score
			}
		}

		return bestScore
	}
}

type Game interface {
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	NextGames() []Game
}

type Player struct {
	Name string
}
