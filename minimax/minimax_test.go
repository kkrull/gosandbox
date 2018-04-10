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
	It("returns a score", func() {
		game := FakeGame{isOver: true}
		score := Minimax(game, max)
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})
})

type FakeGame struct {
	isOver bool
}
