package minimax

type Scorer struct {
	GameStatesFromNextMove interface{
		GameStatesFromNextMove(game Game) []Game
	}
}

func (scorer *Scorer) Score(game Game, player Player) int {
	if player == game.MaximizingPlayer() {
		return scorer.Negamax(game, 1)
	} else {
		return -scorer.Negamax(game, -1)
	}
}

func (scorer *Scorer) Negamax(game Game, polarity int) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, nextGame := range scorer.GameStatesFromNextMove.GameStatesFromNextMove(game) {
		score := -scorer.Negamax(nextGame, -1 * polarity)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

type Game interface {
	AvailableMoves() []Move
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	Move(move Move) Game
}

type Move struct {
	Id string
}

type Player struct {
	Name string
}
