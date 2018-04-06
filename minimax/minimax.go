package minimax

func Minimax(game Game) Outcome {
	if game.FindWinner() == game.MaximizingPlayer() {
		return Outcome{Score: 1}
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return Outcome{Score: -1}
	} else if game.IsOver() {
		return Outcome{Score: 0}
	}

	return Outcome{
		Play: game.AvailablePlays()[0].Play,
	}
}

type Game interface {
	AvailablePlays() []Choice
	FindWinner() Player
	IsOver() bool
	MaximizingPlayer() Player
	MinimizingPlayer() Player
}

type Choice struct {
	Play     Play
	NextGame Game
}

type Play struct {
	Id string
}

type Player struct {
	Name string
}

type Outcome struct {
	Play  Play
	Score int
}
