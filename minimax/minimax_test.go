package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("when the game is over", func() {
		It("returns an error", func() {
			game := CompletedGame()
			_, err := Minimax(game)
			Expect(err).To(MatchError("minimax: game over"))
		})
	})

	Context("when there is only 1 available space", func() {
		It("picks that space", func() {
			game := SingleSpaceGame("Mountain")
			Expect(Minimax(game)).To(Equal("Mountain"))
		})
	})

	Context("when there are 2 of more available spaces", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks that move that causes the maximizing player to win", func() {
				game := CaptureTheFlagGame(EmptySpace("_"), FlagSpace("F"))
				Expect(Minimax(game)).To(Equal("F"))
			})
		})
	})
})


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


/* singleSpaceGame */

func SingleSpaceGame(space Space) Game {
	return singleSpaceGame{openSpaces: []Space{space}}
}

type singleSpaceGame struct {
	openSpaces []Space
}

func (singleSpaceGame) IsOver() bool { return false }
func (g singleSpaceGame) OpenSpaces() []Space { return g.openSpaces }
func (g singleSpaceGame) ClaimSpace(claimed Space) Game {
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
