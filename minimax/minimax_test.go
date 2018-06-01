package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scorer", func() {
	var (
		scorer = &minimax.Scorer{}

		game *GameWithKnownStates
		max  = minimax.Player("Max")
	)

	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game = &GameWithKnownStates{}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})
	})
})

type GameWithKnownStates struct {
}
