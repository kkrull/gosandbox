package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	max = minimax.Player("max")
	min = minimax.Player("min")
)

var _ = Describe("Scorer", func() {
	var (
		scorer = minimax.Scorer{
			MaximizingPlayer: max,
			MinimizingPlayer: min,
		}

		game *GameWithKnownStates
	)

	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game = &GameWithKnownStates{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a game won by the maximizing player as +1", func() {
			game = &GameWithKnownStates{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a game won by the minimizing player as -1", func() {
			game = &GameWithKnownStates{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})

		It("chooses the move with the highest score for the maximizing player", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("Max wins"), &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("chooses the move with the lowest score for the minimizing player", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(minimax.Move("Max loses"), &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		Context("given a game that needs 2 or moves to win", func() {
			BeforeEach(func() {
				game = &GameWithKnownStates{}
				leftTree := &GameWithKnownStates{}
				leftTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				leftTree.AddKnownState(minimax.Move("Max Loses"), &GameWithKnownStates{isOver: true, winner: min})
				game.AddKnownState(minimax.Move("Left"), leftTree)

				rightTree := &GameWithKnownStates{}
				rightTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
				rightTree.AddKnownState(minimax.Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})
				game.AddKnownState(minimax.Move("Right"), rightTree)
			})

			It("the maximizing player assumes the minimizing player will choose the move with the lowest score", func() {
				Expect(scorer.Score(game, max)).To(Equal(0))
			})

			It("the maximizing player assumes the minimizing player will choose the move with the lowest score", func() {
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

func (game *GameWithKnownStates) Move(move minimax.Move) minimax.GameState {
	for i, nextMove := range game.nextMoves {
		if move == nextMove {
			return game.nextGames[i]
		}
	}

	panic("No game found")
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

func (game *GameWithKnownStates) AddKnownState(move minimax.Move, nextGame *GameWithKnownStates) {
	game.nextMoves = append(game.nextMoves, move)
	game.nextGames = append(game.nextGames, nextGame)
}
