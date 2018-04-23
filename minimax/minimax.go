package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, nextGame := range game.GameStatesFromNextMove(player) {
		score := Minimax(nextGame, player)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

type Game interface {
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	GameStatesFromNextMove(player Player) []Game
}

type Player struct {
	Name string
}
