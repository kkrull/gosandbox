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

		It("scores a won by the maximizer as +1", func() {
			game = &GameWithKnownStates{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a won by the minimizer as -1", func() {
			game = &GameWithKnownStates{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})

		It("the maximizer picks the move with the highest score", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("the minimizer picks the move with the lowest score", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("Max Loses"), &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		Context("given a game with 2 or moves left", func() {
			BeforeEach(func() {
				game = &GameWithKnownStates{}
				leftTree := &GameWithKnownStates{}
				game.AddKnownState(minimax.Move("Left"), leftTree)
				leftTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				leftTree.AddKnownState(minimax.Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})

				rightTree := &GameWithKnownStates{}
				game.AddKnownState(minimax.Move("Right"), rightTree)
				rightTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				rightTree.AddKnownState(minimax.Move("Max Loses"), &GameWithKnownStates{isOver: true, winner: min})
			})

			It("the maximizer assumes the minimizer picks the lowest score", func() {
				Expect(scorer.Score(game, max)).To(Equal(0))
			})

			It("the minimizer assumes the maximizer picks the highest score", func() {
				Expect(scorer.Score(game, min)).To(Equal(0))
			})
		})
	})
})

type GameWithKnownStates struct {
	isOver    bool
	winner    minimax.Player
	nextGames []*GameWithKnownStates
	nextMoves []minimax.Move
}

func (game *GameWithKnownStates) AvailableMoves() []minimax.Move {
	return game.nextMoves
}

func (game *GameWithKnownStates) Move(move minimax.Move) minimax.GameState {
	for i, nextMove := range game.nextMoves {
		if nextMove == move {
			return game.nextGames[i]
		}
	}

	panic(move)
}

func (game *GameWithKnownStates) AddKnownState(nextMove minimax.Move, nextGame *GameWithKnownStates) {
	game.nextGames = append(game.nextGames, nextGame)
	game.nextMoves = append(game.nextMoves, nextMove)
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}
