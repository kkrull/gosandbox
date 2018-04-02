package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	maximizer = Player{Name: "Max"}
	minimizer = Player{Name: "Min"}
)

var _ = Describe("Minimax", func() {
	It("returns a result", func() {
		game := FakeGame{}
		Expect(Minimax(game)).To(BeAssignableToTypeOf(Result{}))
	})

	Context("given a game that the maximizing player has won", func() {
		It("scores it as +1", func() {
			game := FakeGame{Winner: maximizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores it as -1", func() {
			game := FakeGame{Winner: minimizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: -1}))
		})
	})
})

type FakeGame struct {
	Winner Player
}

func (game FakeGame) FindWinner() Player {
	return game.Winner
}

func (game FakeGame) MaximizingPlayer() Player {
	return maximizer
}

func (game FakeGame) MinimizingPlayer() Player {
	return minimizer
}
