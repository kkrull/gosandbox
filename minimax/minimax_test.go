package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "Max"}
)

var _ = Describe("Negamax", func() {
	It("scores a game state for a player", func() {
		game := FakeGame{}
		score := Negamax(game, max)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})

	It("returns +1 for a game won by the maximizing player", func() {
		game := FakeGame{Winner: max}
		score := Negamax(game, max)
		Expect(score).To(Equal(1))
	})
})

type FakeGame struct {
	Winner Player
}

