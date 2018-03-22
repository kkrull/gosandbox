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

	Context("when there is only 1 available move", func() {
		It("picks that move", func() {
			game := OneMoveLeftGame("Mountain")
			Expect(Minimax(game)).To(Equal("Mountain"))
		})
	})
})


/* completeGame */

func CompletedGame() Game {
	return completedGame{}
}

type completedGame struct{}
func (completedGame) AvailableSpaces() []Space { return []Space{} }
func (completedGame) IsOver() bool { return true }


/* incompleteGame */

func OneMoveLeftGame(space Space) Game {
	return incompleteGame{OpenSpaces: []Space{space}}
}

type incompleteGame struct {
	OpenSpaces []Space
}

func (g incompleteGame) AvailableSpaces() []Space { return g.OpenSpaces }
func (incompleteGame) IsOver() bool { return false }
