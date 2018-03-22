package minimax

import "fmt"

/* captureTheFlagGame */

func CaptureTheFlagGame(oneBase Space, otherBase Space) Game {
	return captureTheFlagGame{ oneBase, otherBase }
}

type captureTheFlagGame struct {
	oneBase Space
	otherBase Space
}

func (game captureTheFlagGame) ClaimSpace(player Player, claimed Space) Game {
	fmt.Printf("player %v claims %v\n", player, claimed)
	if player == "Max" && claimed == "Min" {
		return victoryGame{Winner: player}
	} else if player == "Min" && claimed == "Max" {
		return victoryGame{Winner: player}
	} else {
		return CompletedGame() //Draw
	}

	panic("Expected somebody to win")
}

func (captureTheFlagGame) IsOver() bool { return false }
func (game captureTheFlagGame) OpenSpaces() []Space {
	return []Space { game.oneBase, game.otherBase }
}

func InstantWinSpace(forPlayer Player) Space {
	return forPlayer
}


/* kingOfTheMountainGame */

func KingOfTheMountainGame(space Space) Game {
	return kingOfTheMountainGame{openSpaces: []Space{space}}
}

type kingOfTheMountainGame struct {
	openSpaces []Space
}

func (kingOfTheMountainGame) IsOver() bool { return false }
func (g kingOfTheMountainGame) OpenSpaces() []Space { return g.openSpaces }
func (g kingOfTheMountainGame) ClaimSpace(Player, Space) Game {
	return completedGame{}
}


/* victoryGame */

type victoryGame struct {
	Winner Player
}

func (victoryGame) ClaimSpace(Player, Space) Game { panic("game over") }
func (victoryGame) IsOver() bool { return true }
func (victoryGame) OpenSpaces() []Space { return []Space{} }


/* completeGame */

func CompletedGame() Game {
	return completedGame{}
}

type completedGame struct{}
func (completedGame) ClaimSpace(Player, Space) Game { panic("game is already over") }
func (completedGame) IsOver() bool { return true }
func (completedGame) OpenSpaces() []Space { return []Space{} }
