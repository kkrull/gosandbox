package minimax

import "fmt"

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
		opponent := opponentOf(player, game)
		nextScoreFromOpponentPerspective := Negamax(nextGame, opponent)
		nextScore := -1 * nextScoreFromOpponentPerspective
		fmt.Printf("Next score: (%s: %d) // (%s: %d)\n", player.Name, nextScore, opponent.Name, nextScoreFromOpponentPerspective)
		if nextScore > maxScorePossible {
			maxScorePossible = nextScore
		}
	}

	return maxScorePossible
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
	NextGames() []Game
}

type Player struct {
	Name string
	Polarity int
}
