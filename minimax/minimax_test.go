package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	max = Player{Name: "Max"}
)

var _ = Describe("Scorer", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			scorer := minimax.Scorer{}
			game := &GameWithKnownStates{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})
	})
})

type GameWithKnownStates struct {
	isOver bool
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

type Player struct {
	Name string
}
