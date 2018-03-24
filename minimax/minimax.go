package minimax

import "fmt"

func Minimax(game Game, player string) (space Space, err error) {
	if game.IsOver() {
		return "", fmt.Errorf("game is already over")
	}

	if player == game.OpponentPlayer() {
		var spaceWithLowestScore Space = nil
		var minScore = 100
		for _, space := range game.AvailableSpaces() {
			nextGame := game.ClaimSpace(space)
			score := nextGame.Score()
			if score < minScore {
				spaceWithLowestScore = space
				minScore = score
			}
		}

		return spaceWithLowestScore, nil
	} else {
		var spaceWithHighestScore Space = nil
		var maxScore = -100
		for _, space := range game.AvailableSpaces() {
			nextGame := game.ClaimSpace(space)
			score := nextGame.Score()
			if score > maxScore {
				spaceWithHighestScore = space
				maxScore = score
			}
		}

		return spaceWithHighestScore, nil
	}
}

type Game interface {
	AvailableSpaces() []Space
	ClaimSpace(Space) Game
	IsOver() bool
	Score() int
	OpponentPlayer() string
}

type Space interface {
}
