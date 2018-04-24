package minimax

type Minimax struct {
	MaximizingPlayer Player
	MinimizingPlayer Player
}

func (scorer Minimax) Score(game Game, player Player) int {
	if game.FindWinner() == scorer.MaximizingPlayer {
		return 1
	} else if game.FindWinner() == scorer.MinimizingPlayer {
		return -1
	} else if game.IsOver() {
		return 0
	}

	if player == scorer.MaximizingPlayer {
		maxScore := -100
		for _, move := range game.AvailableMoves() {
			nextGame := game.Move(move)
			nextScore := scorer.Score(nextGame, scorer.MinimizingPlayer)
			if nextScore > maxScore {
				maxScore = nextScore
			}
		}

		return maxScore
	} else {
		minScore := 100
		for _, move := range game.AvailableMoves() {
			nextGame := game.Move(move)
			nextScore := scorer.Score(nextGame, scorer.MaximizingPlayer)
			if nextScore < minScore {
				minScore = nextScore
			}
		}

		return minScore
	}
}

type Game interface {
	AvailableMoves() []Move
	FindWinner() Player
	IsOver() bool
	Move(move Move) Game
}

type Move struct {
	Id string
}

type Player struct {
	Name string
}
