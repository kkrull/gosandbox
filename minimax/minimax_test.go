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
		game := FakeGame{Over: true}
		Expect(Minimax(game)).To(BeAssignableToTypeOf(Result{}))
	})

	Context("given a game that the maximizing player has won", func() {
		It("scores it as +1", func() {
			game := FakeGame{
				Over:   true,
				Winner: maximizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores it as -1", func() {
			game := FakeGame{
				Over:   true,
				Winner: minimizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: -1}))
		})
	})

	Context("given a game that has ended in a draw", func() {
		It("scores it as 0", func() {
			game := FakeGame{Over: true}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 0}))
		})
	})

	Context("given a game that is not over yet", func() {
		It("picks a move", func() {
			availableMove := FakeMove{Id: "maxWins"}
			game := FakeGame{
				NextMove: availableMove,
				Over:     false}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Move: availableMove}))
		})

		XIt("picks move that ends the game", func() {
		})
	})
})

type FakeGame struct {
	Over     bool
	Winner   Player
	NextMove Move
}

func (game FakeGame) AvailableMoves() []Move {
	return []Move{game.NextMove}
}

func (game FakeGame) FindWinner() Player {
	return game.Winner
}

func (game FakeGame) IsOver() bool {
	return game.Over
}

func (game FakeGame) MaximizingPlayer() Player {
	return maximizer
}

func (game FakeGame) MinimizingPlayer() Player {
	return minimizer
}

type FakeMove struct {
	Id string
}
