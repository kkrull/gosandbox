package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("given a game that the maximizing player has won", func() {
		It("scores that game as +1", func() {
			game := CompletedGame{Winner: "Max"}
			Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{Score: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores that game as -1", func() {
			game := CompletedGame{Winner: "Min"}
			Expect(Minimax(game, "Min")).To(BeEquivalentTo(Result{Score: -1}))
		})
	})

	Context("given a game ending in a draw", func() {
		It("scores that game as -1", func() {
			game := CompletedGame{}
			Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{Score: 0}))
		})
	})
	

})

type CompletedGame struct {
	Winner string
}

func (game CompletedGame) FindWinner() string {
	return game.Winner
}

