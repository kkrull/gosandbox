package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else if game.IsOver() {
		return 0
	}

	bestScore := 100
	for _, nextGame := range game.NextGames() {
		nextScore := Minimax(nextGame, Player{})
		if nextScore < bestScore {
			bestScore = nextScore
		}
	}

	return bestScore
}

type Game interface {
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	NextGames() []Game
}

type Player struct {
	Name string
}
