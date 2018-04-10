package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "max", Polarity: 1}
	min = Player{Name: "min", Polarity: -1}
)

var _ = Describe("Negamax", func() {
	It("scores a game ending in a draw as 0", func() {
		game := FakeGame{isOver: true}
		Expect(Negamax(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as +1", func() {
		game := FakeGame{isOver: true, winner: max}
		Expect(Negamax(game, max)).To(Equal(1))
	})

	It("scores a game won by the minimizing player as -1", func() {
		game := FakeGame{isOver: true, winner: min}
		Expect(Negamax(game, max)).To(Equal(-1))
	})

	It("scores a game won by the maximizing player as -1, from the minimizing player's perspective", func() {
		game := FakeGame{isOver: true, winner: max}
		Expect(Negamax(game, min)).To(Equal(-1))
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
