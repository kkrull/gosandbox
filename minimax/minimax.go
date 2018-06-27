package minimax

type Scorer struct {
	Maximizer Player
	Minimizer Player
}

func (scorer *Scorer) Score(game GameState, player Player) int {
	if game.FindWinner() == scorer.Maximizer {
		return 1
	} else if game.FindWinner() == scorer.Minimizer {
		return -1
	} else if game.IsOver() {
		return 0
	}

	nextMove := game.AvailableMoves()[0]
	nextGame := game.Move(nextMove)
	return scorer.Score(nextGame, Player("Who?"))
}

type GameState interface {
	AvailableMoves() []Move
	FindWinner() Player
	IsOver() bool
	Move(move Move) GameState
}

type Move string
type Player string
