package minimax

func Minimax(game Game, player Player) int {
	if game.FindWinner() == game.MaximizingPlayer() {
		return 1
	} else if game.FindWinner() == game.MinimizingPlayer() {
		return -1
	} else if game.IsOver() {
		return 0
	}

	if player == game.MaximizingPlayer() {
		maxScore := -100
		for _, nextGame := range game.GameStatesFromNextMove(player) {
			score := Minimax(nextGame, game.MinimizingPlayer())
			if score > maxScore {
				maxScore = score
			}
		}

		return maxScore
	} else {
		minScore := 100
		for _, nextGame := range GameStatesFromNextMove(game, player) {
			score := Minimax(nextGame, game.MaximizingPlayer())
			if score < minScore {
				minScore = score
			}
		}

		return minScore
	}
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
	IsOver() bool
	FindWinner() Player
	MaximizingPlayer() Player
	MinimizingPlayer() Player
	AvailableMoves(player Player) []Move
	Move(player Player, move Move) Game
	//GameStatesFromNextMove(player Player) []Game
}

type Move struct {
	Id string
}

type Player struct {
	Name string
}
