package minimax

func Minimax(game Game, player Player) int {
	if player == game.MaximizingPlayer() {
		return Negamax(game, 1)
	} else {
		return -Negamax(game, -1)
	}
}

func Negamax(game Game, polarity int) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1 * polarity
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1 * polarity
	} else if game.IsOver() {
		return 0
	}

	maxScore := -100
	for _, nextGame := range GameStatesFromNextMove(game) {
		score := -Negamax(nextGame, -1 * polarity)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func GameStatesFromNextMove(game Game) []Game {
	nextGames := make([]Game, 0)
	for _, move := range game.AvailableMoves() {
		nextGame := game.Move(move)
		nextGames = append(nextGames, nextGame)
	}

	return nextGames
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
