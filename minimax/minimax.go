package minimax

func Negamax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * player.Polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * player.Polarity
	} else if game.IsOver() {
		return 0
	}

	maxScorePossible := -100
	for _, nextGame := range game.NextGames() {
		opponent := player
		nextScore := -Negamax(nextGame, opponent)
		//fmt.Printf("Next score from %s: %d\n", opponent.Name, nextScore)
		if nextScore > maxScorePossible {
			maxScorePossible = nextScore
		}
	}

	return maxScorePossible
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
	Polarity int
}
