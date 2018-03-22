package minimax

/* captureTheFlagGame */

func CaptureTheFlagGame(emptySpace Space, flagSpace Space) Game {
	return captureTheFlagGame{
		flagSpace: flagSpace,
		openSpaces: []Space{emptySpace, flagSpace},
	}
}

type captureTheFlagGame struct {
	flagSpace Space
	openSpaces []Space
}

func (game captureTheFlagGame) ClaimSpace(claimed Space) Game {
	if claimed == game.flagSpace {
		return completedGame{}
	}

	var unclaimedSpaces []Space
	for _, openSpace := range game.openSpaces {
		if openSpace != claimed {
			unclaimedSpaces = append(unclaimedSpaces, openSpace)
		}
	}

	return captureTheFlagGame{
		flagSpace: game.flagSpace,
		openSpaces: unclaimedSpaces,
	}
}

func (captureTheFlagGame) IsOver() bool { return false }
func (game captureTheFlagGame) OpenSpaces() []Space { return game.openSpaces }

func EmptySpace(id string) Space {
	return id
}

func FlagSpace(id string) Space {
	return id
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
func (g kingOfTheMountainGame) ClaimSpace(claimed Space) Game {
	return completedGame{}
}


/* completeGame */

func CompletedGame() Game {
	return completedGame{}
}

type completedGame struct{}
func (completedGame) ClaimSpace(Space) Game { panic("game is already over") }
func (completedGame) IsOver() bool { return true }
func (completedGame) OpenSpaces() []Space { return []Space{} }
