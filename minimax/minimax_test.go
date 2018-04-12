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
	It("scores a draw game as 0", func() {
		game := FakeGame{over: true}
		Expect(Minimax(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as +1", func() {
		game := FakeGame{over: true, winner: max}
		Expect(Minimax(game, max)).To(Equal(1))
	})

	It("scores a game won by the minimizing player as -1", func() {
		game := FakeGame{over: true, winner: min}
		Expect(Minimax(game, max)).To(Equal(-1))
	})

	It("picks the maximum score the maximizing player can achieve, assuming the minimizing player picks the move with the lowest score", func() {
		game := FakeGame{nextGames: []Game{
			FakeGame{over: true},
			FakeGame{over: true, winner: min}}}
		Expect(Minimax(game, max)).To(Equal(0))
	})
})

type FakeGame struct {
	nextGames []Game
	over      bool
	winner    Player
}

func (game FakeGame) FindWinner() Player {
	return game.winner
}

func (game FakeGame) IsOver() bool {
	return game.over
}

func (game FakeGame) MaximizingPlayer() Player {
	return max
}

func (game FakeGame) MinimizingPlayer() Player {
	return min
}

func (game FakeGame) NextGames() []Game {
	return game.nextGames
}
