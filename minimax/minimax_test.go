package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	It("returns a result", func() {
		game := FakeGame{}
		Expect(Minimax(game)).To(BeAssignableToTypeOf(Result{}))
	})

	Context("given a game that the maximizing player has won", func() {
		It("scores it as +1", func() {
			game := FakeGame{Winner: "Max"}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 1}))
		})
	})
})

type FakeGame struct {
	Winner string
}
