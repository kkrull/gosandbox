package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "max"}
)

var _ = Describe("Minimax", func() {
	It("scores a draw game as 0", func() {
		game := FakeGame{over: true}
		Expect(Minimax(game, max)).To(Equal(0))
	})

	XIt("scores a game won by the maximizing player as +1", func() {
		game := FakeGame{over: true, winner: max}
		Expect(Minimax(game, max)).To(Equal(1))
	})
})

type FakeGame struct {
	over bool
	winner Player
}

func (game FakeGame) IsOver() bool {
	return game.over
}

