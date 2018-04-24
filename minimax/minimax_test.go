package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "Max"}
)

var _ = Describe("Minimax", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game := FakeGame{isOver: true}
			scorer := Minimax{}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})
	})
})

type FakeGame struct {
	isOver bool
}
