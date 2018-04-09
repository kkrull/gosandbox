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
		game := FakeGame{isOver: true}
		score := Negamax(game, max)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})

	It("returns +1 for a game won by the maximizing player", func() {
		game := FakeGame{isOver: true, winner: max}
		score := Negamax(game, max)
		Expect(score).To(Equal(1))
	})

	It("returns -1 for a game won by the minimizing player", func() {
		game := FakeGame{isOver: true, winner: min}
		score := Negamax(game, max)
		Expect(score).To(Equal(-1))
	})

	It("returns 0 for a game ending in a draw", func() {
		game := FakeGame{isOver: true}
		score := Negamax(game, max)
		Expect(score).To(Equal(0))
	})

	It("scores a game that is not over", func() {
		game := FakeGame{successors: []Game{
			FakeGame{isOver: true, winner: max},
		}}

		score := Negamax(game, max)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})

	It("maximizes the score when it is the maximizing player's turn", func() {
		game := FakeGame{successors: []Game{
			FakeGame{isOver: true, winner: min},
			FakeGame{isOver: true, winner: max},
		}}

		score := Negamax(game, max)
		Expect(score).To(Equal(1))
	})

	It("maximizes the minimum score achieved by the opponent, when it is the minimizing player's turn", func() {
		game := FakeGame{successors: []Game{
			FakeGame{isOver: true, winner: max}, //-1
			FakeGame{isOver: true, winner: min}, //+1
		}}

		score := Negamax(game, min)
		Expect(score).To(Equal(1))
	})
})

type FakeGame struct {
	isOver     bool
	winner     Player
	successors []Game
}

func (game FakeGame) FindWinner() Player {
	return game.winner
}

func (game FakeGame) IsOver() bool {
	return game.isOver
}

func (FakeGame) MaximizingPlayer() Player {
	return max
}

func (FakeGame) MinimizingPlayer() Player {
	return min
}

func (game FakeGame) Successors() []Game {
	return game.successors
}
