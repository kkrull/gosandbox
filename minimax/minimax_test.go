package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "Max"}
	min = Player{Name: "Min"}
)

var _ = Describe("Negamax", func() {
	It("scores a game state for a player", func() {
		game := FakeGame{Over: true}
		score := Negamax(game, max)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})

	It("returns +1 for a game won by the maximizing player", func() {
		game := FakeGame{Over: true, Winner: max}
		score := Negamax(game, max)
		Expect(score).To(Equal(1))
	})

	It("returns -1 for a game won by the minimizing player", func() {
		game := FakeGame{Over: true, Winner: min}
		score := Negamax(game, max)
		Expect(score).To(Equal(-1))
	})

	It("returns 0 for a game ending in a draw", func() {
		game := FakeGame{Over: true}
		score := Negamax(game, max)
		Expect(score).To(Equal(0))
	})
})

type FakeGame struct {
	Over bool
	Winner Player
}

func (game FakeGame) FindWinner() Player {
	return game.Winner
}

func (game FakeGame) IsOver() bool {
	return game.Over
}

func (FakeGame) MaximizingPlayer() Player {
	return max
}

func (FakeGame) MinimizingPlayer() Player {
	return min
}
