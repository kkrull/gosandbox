package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	max    = Player{Name: "Max"}
	min    = Player{Name: "Min"}
	scorer = minimax.Scorer{
		MaximizingPlayer: max,
		MinimizingPlayer: min,
	}
)

var _ = Describe("Scorer", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game := &GameWithKnownStates{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a game won by the maximizing player as +1", func() {
			game := &GameWithKnownStates{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a game won by the minimizing player as -1", func() {
			game := &GameWithKnownStates{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})
	})
})

type GameWithKnownStates struct {
	isOver bool
	winner Player
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

type Player struct {
	Name string
}
