package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Negamax", func() {
	It("scores a game state for a player", func() {
		game := FakeGame{}
		player := Player{Name: "Max"}
		score := Negamax(game, player)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})
})

type FakeGame struct {
}
