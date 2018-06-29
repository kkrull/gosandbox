package minimax

func NewComputerPlayer(opponent Player) *ComputerPlayer {
	scorer := &Scorer{Minimizer: opponent}
	computerPlayer := &ComputerPlayer{
		Scorer:   scorer,
		Opponent: opponent,
	}
	scorer.Maximizer = computerPlayer
	return computerPlayer
}

type ComputerPlayer struct {
	Scorer   GameScorer
	Opponent Player
}

func (player *ComputerPlayer) PickMove(game GameState) Move {
	var best *scoredMove
	for _, nextMove := range game.AvailableMoves() {
		nextGame := game.Move(nextMove)
		nextScore := player.Scorer.Score(nextGame, player.Opponent)
		if best == nil || nextScore > best.score {
			best = &scoredMove{move: nextMove, score: nextScore}
		}
	}

	return best.move
}

type GameScorer interface {
	Score(game GameState, player Player) int
}

type scoredMove struct {
	move  Move
	score int
}
