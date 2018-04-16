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
	It("scores a game ending in a draw as 0", func() {
		game := FakeGame{}
		Expect(Minimax(game, max)).To(Equal(0))
	})
})

type FakeGame struct {

}
