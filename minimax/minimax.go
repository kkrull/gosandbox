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
		maxScore := -100
		for _, nextGame := range game.NextGames() {
			nextScore := Minimax(nextGame, game.MinimizingPlayer())
			if nextScore > maxScore {
				maxScore = nextScore
			}
		}

		return maxScore
	} else {
		minScore := 100
		for _, nextGame := range game.NextGames() {
			nextScore := Minimax(nextGame, game.MaximizingPlayer())
			if nextScore < minScore {
				minScore = nextScore
			}
		}

		return minScore
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
