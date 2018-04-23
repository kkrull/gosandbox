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
	for _, nextGame := range GameStatesFromNextMove(game, game.MaximizingPlayer()) {
		score := -Negamax(nextGame, -1 * polarity)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func GameStatesFromNextMove(game Game, player Player) []Game {
	nextGames := make([]Game, 0)
	for _, move := range game.AvailableMoves(player) {
		nextGame := game.Move(player, move)
		nextGames = append(nextGames, nextGame)
	}

	return nextGames
}

type Game interface {
	AvailableMoves(player Player) []Move
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	Move(player Player, move Move) Game
}

type Move struct {
	Id string
}

type Player struct {
	Name string
}
