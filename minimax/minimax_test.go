package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max  = Player{Name: "Max"}
	game Game
)

var _ = Describe("Minimax", func() {
	It("scores a game ending in a draw as 0", func() {
		game = &FakeGame{isOver: true}
		Expect(Minimax(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as 1", func() {
		game = &FakeGame{isOver: true, winner: max}
		Expect(Minimax(game, max)).To(Equal(1))
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

func (game *FakeGame) MaximizingPlayer() Player {
	return max
}
