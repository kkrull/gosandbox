package minimax

import "fmt"

func Minimax(game Game, player Player) Outcome {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Outcome{BoundedScore: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Outcome{BoundedScore: -1}
	} else if game.IsOver() {
		return Outcome{BoundedScore: 0}
	}

	suspiciouslyAmbiguousOpponent := Player{}
	if player == game.MaximizingPlayer() {
		bestOutcome := Outcome{BoundedScore: -100}
		for _, choice := range game.AvailableChoices() {
			outcome := Minimax(choice.ResultingGame, suspiciouslyAmbiguousOpponent)
			if outcome.BoundedScore > bestOutcome.BoundedScore {
				bestOutcome = Outcome{
					BoundedScore: outcome.BoundedScore,
					Play:         choice.Play}
			}
		}

		return bestOutcome
	} else if player == game.MinimizingPlayer() {
		bestOutcome := Outcome{BoundedScore: 100}
		for _, choice := range game.AvailableChoices() {
			outcome := Minimax(choice.ResultingGame, suspiciouslyAmbiguousOpponent)
			if outcome.BoundedScore < bestOutcome.BoundedScore {
				bestOutcome = Outcome{
					BoundedScore: outcome.BoundedScore,
					Play:         choice.Play}
			}
		}

		return bestOutcome
	} else {
		panic(fmt.Sprintf("unknown player: %s", player.Name))
	}
}

type Game interface {
	AvailableChoices() []Choice
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Choice struct {
	Play          Play
	ResultingGame Game
}

type Player struct {
	Name string
}

type Outcome struct {
	Play         Play
	BoundedScore int
}

type Play interface{}
