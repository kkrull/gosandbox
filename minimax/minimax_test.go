package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scorer", func() {
	var (
		game   *GameWithKnownStates
		max    = minimax.Player("Max")
		min    = minimax.Player("Min")
		scorer = &minimax.Scorer{
			Maximizer: max,
			Minimizer: min,
		}
	)

	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game = &GameWithKnownStates{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a game won by the maximizer as +1", func() {
			game = &GameWithKnownStates{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a game won by the minimizer as -1", func() {
			game = &GameWithKnownStates{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})

		It("the maximizer picks the move with the highest score in a game that is not over", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("MaxWins"), &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("the minimizer picks the move with the lowest score in a game that is not over", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("MaxLoses"), &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		Context("given a game that has 2 or moves left", func() {
			BeforeEach(func() {
				game = &GameWithKnownStates{}
				leftTree := &GameWithKnownStates{}
				leftTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				leftTree.AddKnownState(minimax.Move("MaxWins"), &GameWithKnownStates{isOver: true, winner: max})
				game.AddKnownState(minimax.Move("Left"), leftTree)

				rightTree := &GameWithKnownStates{}
				rightTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				rightTree.AddKnownState(minimax.Move("MaxLoses"), &GameWithKnownStates{isOver: true, winner: min})
				game.AddKnownState(minimax.Move("Right"), rightTree)
			})

			It("the maximizer assumes the minimizer picks the move with the lowest score", func() {
				Expect(scorer.Score(game, max)).To(Equal(0))
			})
			It("the minimizer assumes the maximizer picks the move with the highest score", func() {
				Expect(scorer.Score(game, min)).To(Equal(0))
			})
		})
	})
})

type GameWithKnownStates struct {
	isOver    bool
	winner    minimax.Player
	nextMoves []minimax.Move
	nextGames []*GameWithKnownStates
}

func (game *GameWithKnownStates) AvailableMoves() []minimax.Move {
	return game.nextMoves
}

func (game *GameWithKnownStates) AddKnownState(move minimax.Move, nextGame *GameWithKnownStates) {
	game.nextMoves = append(game.nextMoves, move)
	game.nextGames = append(game.nextGames, nextGame)
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

func (game *GameWithKnownStates) Move(move minimax.Move) minimax.GameState {
	for i, nextMove := range game.nextMoves {
		if nextMove == move {
			return game.nextGames[i]
		}
	}

	panic("move not found")
}
