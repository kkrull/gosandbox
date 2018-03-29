package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("given a game that the maximizing player has won", func() {
		It("scores that game as +1", func() {
			game := VictoryGame{Winner: "Max"}
			Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{Score: 1}))
		})
	})
})

type VictoryGame struct {
	Winner string
}
