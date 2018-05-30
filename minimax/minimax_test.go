package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	max    = Player("Max")
	min    = Player("Min")
	scorer *minimax.Scorer
	game   *GameWithKnownStates
)

var _ = Describe("Scorer", func() {
	Describe("#Score", func() {
		BeforeEach(func() {
			scorer = &minimax.Scorer{
				MaximizingPlayer: max,
				MinimizingPlayer: min,
			}
		})

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

		It("picks the maximum score for the maximizing player", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("picks the minimum score for the minimum player", func() {
			game = &GameWithKnownStates{}
			game.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
			game.AddKnownState(Move("Min Wins"), &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})

		Context("given an unfinished game with 2 or moves left", func() {
			BeforeEach(func() {
				game = &GameWithKnownStates{}
				leftFork := &GameWithKnownStates{}
				leftFork.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
				leftFork.AddKnownState(Move("Max Wins"), &GameWithKnownStates{isOver: true, winner: max})
				game.AddKnownState(Move("Left"), leftFork)

				rightFork := &GameWithKnownStates{}
				rightFork.AddKnownState(Move("Draw"), &GameWithKnownStates{isOver: true})
				rightFork.AddKnownState(Move("Min Wins"), &GameWithKnownStates{isOver: true, winner: min})
				game.AddKnownState(Move("Right"), rightFork)
			})

			It("the maximizing player assumes the minimizing player will pick the move with the lowest score", func() {
				Expect(scorer.Score(game, max)).To(Equal(0))
			})

			It("the minimizing player assumes the maximizing player will pick the move with the highest score", func() {
				Expect(scorer.Score(game, min)).To(Equal(0))
			})
		})
	})
})

type GameWithKnownStates struct {
	isOver    bool
	winner    Player
	moves     []Move
	nextGames []*GameWithKnownStates
}

func (game *GameWithKnownStates) AddKnownState(move Move, nextGame *GameWithKnownStates) {
	game.moves = append(game.moves, move)
	game.nextGames = append(game.nextGames, nextGame)
}

func (game *GameWithKnownStates) AvailableMoves() []minimax.Move {
	moves := make([]minimax.Move, len(game.moves))
	for i, move := range game.moves {
		moves[i] = move
	}

	return moves
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) Move(move minimax.Move) minimax.GameState {
	for i, availableMove := range game.moves {
		if move == availableMove {
			return game.nextGames[i]
		}
	}

	panic("unknown move")
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

type Move string

type Player string
