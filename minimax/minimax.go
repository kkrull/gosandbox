package minimax

type Minimax struct {

}

func (scorer Minimax) Score(game Game, player Player) int {
	return 0
}

type Game interface {

}

type Player struct {
	Name string
}
