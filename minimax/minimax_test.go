package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max    = Player{Name: "Max"}
	scorer = Minimax{
		MaximizingPlayer: max,
	}
)

var _ = Describe("Minimax", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game := &FakeGame{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a gamw won by the maximizing player as +1", func() {
			game := &FakeGame{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
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
