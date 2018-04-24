package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max    = Player{Name: "Max"}
	min    = Player{Name: "Min"}
	scorer = Minimax{
		MaximizingPlayer: max,
		MinimizingPlayer: min,
	}
)

var _ = Describe("Minimax", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game := &FakeGame{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a game won by the maximizing player as +1", func() {
			game := &FakeGame{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a game won by the minimizing player as -1", func() {
			game := &FakeGame{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})
	})
})

type FakeGame struct {
	isOver bool
	winner Player
}

func (game *FakeGame) FindWinner() Player {
	return game.winner
}

func (game *FakeGame) IsOver() bool {
	return game.isOver
}
