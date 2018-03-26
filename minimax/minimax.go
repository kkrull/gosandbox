package minimax

import (
	"fmt"
)

func Minimax(game Game) (result Result, err error) {
	if game.IsOver() {
		return Result{Space: ""}, fmt.Errorf("the game is already over")
	}

	//var bestScore int = math.Inf(-1)
	//for _, space := range game.OpenSpaces() {
	//	nextGame := game.ClaimSpace(space)
	//	result, _ := Minimax(nextGame)
	//
	//	if result.Score > bestScore {
	//		bestScore = result.Score
	//	}
	//}

	space := game.OpenSpaces()[0]
	return Result{Space: space}, nil
	//return 0, fmt.Errorf("unable to pick a space")
}

type Game interface {
	OpenSpaces() []Space
	IsOver() bool
}

type Result struct {
	Space Space
}

type Space interface {
}
