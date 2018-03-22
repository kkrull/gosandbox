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
			game := OneSpaceLeftGame("Mountain")
			Expect(Minimax(game)).To(Equal("Mountain"))
		})
	})

	Context("when there are 2 of more available spaces", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks that move that causes the maximizing player to win", func() {
				game := MultiSpaceGame(EmptySpace("_"), FlagSpace("F"))
				Expect(Minimax(game)).To(Equal("F"))
			})
		})
	})
})


/* completeGame */

func CompletedGame() Game {
	return completedGame{}
}

type completedGame struct{}
func (completedGame) ClaimSpace(Space) Game { panic("game is already over") }
func (completedGame) IsOver() bool { return true }
func (completedGame) OpenSpaces() []Space { return []Space{} }


/* incompleteGame */

func MultiSpaceGame(spaces ...Space) Game {
	return incompleteGame{openSpaces: spaces}
}

func OneSpaceLeftGame(space Space) Game {
	return incompleteGame{openSpaces: []Space{space}}
}

type incompleteGame struct {
	openSpaces []Space
}

func (incompleteGame) IsOver() bool { return false }
func (g incompleteGame) OpenSpaces() []Space { return g.openSpaces }
func (g incompleteGame) ClaimSpace(space Space) Game {
	var unclaimedSpaces []Space
	for _, openSpace := range g.openSpaces {
		if openSpace != openSpace {
			unclaimedSpaces = append(unclaimedSpaces, space)
		}
	}

	return incompleteGame{openSpaces: unclaimedSpaces}
}


/* Spaces */

func EmptySpace(id string) Space {
	return id
}

func FlagSpace(id string) Space {
	return id
}
