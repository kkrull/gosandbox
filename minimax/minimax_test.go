package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scorer", func() {
	var (
		scorer *minimax.Scorer
		game   *GameWithKnownStates
	)

	Describe("#Score", func() {
		BeforeEach(func() {
			scorer = &minimax.Scorer{}
		})

		It("scores a game ending in a draw as 0", func() {
			game = &GameWithKnownStates{}
			Expect("Universal Answer").To(Equal(42))
		})
	})
})

type GameWithKnownStates struct {
}
