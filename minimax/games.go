package minimax

import (
	"fmt"
)

/* EndGame */

func EndGame() Game {
	return endGame{}
}

type endGame struct{}

func (endGame) ClaimedSpaces() map[string]string {
	panic("implement me")
}

func (endGame) OpponentPlayer() string {
	return "Opponent"
}

func (endGame) Score() int {
	return 0 //TODO KDK: Did anyone win?
}

func (endGame) ClaimSpace(Space) Game {
	panic("game is already over")
}

func (endGame) AvailableSpaces() []Space {
	return []Space{}
}

func (endGame) IsOver() bool {
	return true
}

/* OpenGame */

func OpenGame(availableSpaces ...Space) Game {
	return openGame{
		playerName:      "Player",
		opponentName:    "Opponent",
		availableSpaces: availableSpaces,
	}
}

type openGame struct {
	playerName      string
	opponentName    string
	availableSpaces []Space
}

func (game openGame) ClaimedSpaces() map[string]string {
	panic("implement me")
}

func (game openGame) OpponentPlayer() string {
	return game.opponentName
}

func (game openGame) ClaimSpace(claimed Space) Game {
	spaces := removeSpace(claimed, game.AvailableSpaces())
	return makeGame(spaces)
}

func (game openGame) Score() int {
	var claims map[string]string = game.ClaimedSpaces()
	if claims["PlayerWins"] == "Player" {
		return 1
	} else if claims["OpponentWins"] == "Opponent" {
		return -1
	}

	//switch game.availableSpaces[0] {
	//case "OpponentWins":
	//	return 1
	//case "PlayerWins":
	//	return -1
	//case "Useless":
	//	return 1
	//default:
	//panic(fmt.Sprintf("Unknown score for remaining space %v", game.availableSpaces))
	//}

	panic(fmt.Sprintf("Unknown score for remaining space %v", game.availableSpaces))
}

func (game openGame) AvailableSpaces() []Space {
	return game.availableSpaces
}

func (openGame) IsOver() bool {
	return false
}

func makeGame(spaces []Space) Game {
	if len(spaces) == 0 {
		return EndGame()
	}

	return openGame{availableSpaces: spaces}
}

/* TwoStepGame */

func TwoStepGame(firstStepSpace Space, uselessSpaces ...Space) Game {
	return twoStepGame{
		firstStepSpace: firstStepSpace,
		uselessSpaces:  uselessSpaces,
	}
}

type twoStepGame struct {
	firstStepSpace Space
	uselessSpaces  []Space
}

func (game twoStepGame) AvailableSpaces() []Space {
	return append(game.uselessSpaces, game.firstStepSpace)
}

func (game twoStepGame) ClaimSpace(claimed Space) Game {
	if claimed == "PlayerStep1" {
		return OpenGame(append(game.uselessSpaces, "PlayerWins"))
	}

	spaces := removeSpace(claimed, game.AvailableSpaces())
	return makeGame(spaces)
}

func (game twoStepGame) IsOver() bool {
	return false
}

func (twoStepGame) Score() int {
	panic(fmt.Sprintf("Figure out how to score"))
}

func (twoStepGame) OpponentPlayer() string {
	return "Opponent"
}

/* Helpers */

func removeSpace(remove Space, fromSpaces []Space) []Space {
	var output []Space
	for _, space := range fromSpaces {
		if space != remove {
			output = append(output, space)
		}
	}

	return output
}
