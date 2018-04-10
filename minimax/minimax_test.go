package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "max"}
	min = Player{Name: "min"}
)

var _ = Describe("Minimax", func() {
	It("scores a game ending in a draw as 0", func() {
		game := FakeGame{isOver: true}
		Expect(Minimax(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as +1", func() {
		game := FakeGame{isOver: true, winner: max}
		Expect(Minimax(game, max)).To(Equal(1))
	})

	It("scores a game won by the minimizing player as -1", func() {
		game := FakeGame{isOver: true, winner: min}
		Expect(Minimax(game, max)).To(Equal(-1))
	})
})

type FakeGame struct {
	isOver bool
	winner Player
}

func (game FakeGame) FindWinner() Player {
	return game.winner
}

func (game FakeGame) MaximizingPlayer() Player {
	return max
}

func (game FakeGame) MinimizingPlayer() Player {
	return min
}
