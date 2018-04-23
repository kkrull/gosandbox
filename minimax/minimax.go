package minimax

func Minimax(game Game, player Player) int {
	if game.IsOver() {
		return 0
	}

	return 999
}

type Game interface {
	IsOver() bool
}

type Player struct {
	Name string
}
