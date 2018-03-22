package minimax

/* captureTheFlagGame */

func CaptureTheFlagGame(oneBase Space, otherBase Space) Game {
	return captureTheFlagGame{ oneBase, otherBase }
}

type captureTheFlagGame struct {
	oneBase Space
	otherBase Space
}

func (game captureTheFlagGame) ClaimSpace(player Player, claimed Space) Game {
	if player == "Max" && claimed == "MinBase" {
		return victoryGame{Winner: player}
	} else if player == "Min" && claimed == "MaxBase" {
		return victoryGame{Winner: player}
	} else {
		return DrawGame()
	}

	panic("Expected somebody to win")
}

func (captureTheFlagGame) FindWinner() Player { return nil }
func (captureTheFlagGame) IsOver() bool { return false }
func (game captureTheFlagGame) OpenSpaces() []Space {
	return []Space { game.oneBase, game.otherBase }
}

func HomeBaseSpace(ofPlayer Player) Space {
	return ofPlayer
}


/* kingOfTheMountainGame */

func KingOfTheMountainGame(space Space) Game {
	return kingOfTheMountainGame{openSpaces: []Space{space}}
}

type kingOfTheMountainGame struct {
	openSpaces []Space
}

func (kingOfTheMountainGame) IsOver() bool { return false }
func (kingOfTheMountainGame) FindWinner() Player { return nil }
func (g kingOfTheMountainGame) OpenSpaces() []Space { return g.openSpaces }
func (g kingOfTheMountainGame) ClaimSpace(player Player, space Space) Game {
	return victoryGame{Winner: player}
}


/* victoryGame */

type victoryGame struct {
	Winner Player
}

func (victoryGame) ClaimSpace(Player, Space) Game { panic("game over") }
func (g victoryGame) FindWinner() Player { return g.Winner }
func (victoryGame) IsOver() bool { return true }
func (victoryGame) OpenSpaces() []Space { return []Space{} }


/* drawGame */

func DrawGame() Game {
	return drawGame{}
}

type drawGame struct{}
func (drawGame) ClaimSpace(Player, Space) Game { panic("game is already over") }
func (drawGame) FindWinner() Player { return nil }
func (drawGame) IsOver() bool { return true }
func (drawGame) OpenSpaces() []Space { return []Space{} }
