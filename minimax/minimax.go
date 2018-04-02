package minimax

func Minimax(game Game) Outcome {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Outcome{BoundedScore: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Outcome{BoundedScore: -1}
	} else if game.IsOver() {
		return Outcome{BoundedScore: 0}
	}

	bestOutcome := Outcome{BoundedScore: -100}
	for _, choice := range game.AvailableChoices() {
		outcome := Minimax(choice.ResultingGame)
		if outcome.BoundedScore > bestOutcome.BoundedScore {
			bestOutcome = Outcome{
				BoundedScore: outcome.BoundedScore,
				Play:         choice.Play}
		}
	}

	return bestOutcome
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
