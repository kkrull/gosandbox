package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	game *GameWithKnownStates
	max  = minimax.Player("Max")
	min  = minimax.Player("Min")
)

var _ = Describe("Scorer", func() {
	var scorer *minimax.Scorer
	BeforeEach(func() {
		scorer = &minimax.Scorer{
			MaximizingPlayer: max,
			MinimizingPlayer: min,
		}
	})

	Describe("#Score", func() {
		It("returns 0 for a game ending in a draw", func() {
			game = &GameWithKnownStates{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("returns +1 for a game won by the maximizing player", func() {
			game = &GameWithKnownStates{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("returns -1 for a game won by the minimizing player", func() {
			game = &GameWithKnownStates{isOver: true, winner: min}
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		It("returns the maximum possible score when the maximizing player has more than 1 choice", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("returns the minimum possible score when the minimizing player has more than 1 choice", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(Move("Max Loses"), &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		Context("given an unfinished game", func() {
			BeforeEach(func() {
				game = &GameWithKnownStates{}
				leftGameTree := &GameWithKnownStates{}
				leftGameTree.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true}) //min picks
				leftGameTree.AddKnownState(Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max}) //max picks
				game.AddKnownState(Move("Left"), leftGameTree)

				rightGameTree := &GameWithKnownStates{}
				rightGameTree.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true}) //max picks
				rightGameTree.AddKnownState(Move("Max Loses"), &GameWithKnownStates{isOver: true, winner: min}) //min picks
				game.AddKnownState(Move("Right"), rightGameTree)
			})

			It("the maximizing player assumes the minimizing player will minimize the score", func() {
				Expect(scorer.Score(game, max)).To(Equal(0))
			})

			It("the minimizing player assumes the maximizing player will maximize the score", func() {
				Expect(scorer.Score(game, min)).To(Equal(0))
			})
		})
	})
})

type GameWithKnownStates struct {
	isOver         bool
	winner         minimax.Player
	moves          []Move
	nextGameStates []*GameWithKnownStates
}

func (game *GameWithKnownStates) AddKnownState(move Move, nextGameState *GameWithKnownStates) {
	game.moves = append(game.moves, move)
	game.nextGameStates = append(game.nextGameStates, nextGameState)
}

func (game *GameWithKnownStates) AvailableMoves() []minimax.Move {
	moves := make([]minimax.Move, len(game.moves))
	for i, move := range game.moves {
		moves[i] = move
	}

	return moves
}

func (game *GameWithKnownStates) NextGameState(move minimax.Move) minimax.GameState {
	for i, availableMove := range game.moves {
		if availableMove == move {
			return game.nextGameStates[i]
		}
	}

	panic("Invalid move")
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

type Move string
