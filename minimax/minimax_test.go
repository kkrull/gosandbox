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
func (completedGame) IsOver() bool { return true }
func (completedGame) OpenSpaces() []Space { return []Space{} }


/* incompleteGame */

func OneMoveLeftGame(space Space) Game {
	return incompleteGame{openSpaces: []Space{space}}
}

type incompleteGame struct {
	openSpaces []Space
}

func (incompleteGame) IsOver() bool { return false }
func (g incompleteGame) OpenSpaces() []Space { return g.openSpaces }
