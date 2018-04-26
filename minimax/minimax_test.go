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

		It("scores the maximum possible score for the maximizing player in an unfinished game", func() {
			game := &GameWithKnownStates{}
			game.AddStateTransition(Move{Id: "Max chooses...poorly"}, &GameWithKnownStates{isOver: true})
			game.AddStateTransition(Move{Id: "Max chooses...wisely"}, &GameWithKnownStates{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores the minimum possible score for the minimizing player in an unfinished game", func() {
			game := &GameWithKnownStates{}
			game.AddStateTransition(Move{Id: "Min chooses...poorly"}, &GameWithKnownStates{isOver: true})
			game.AddStateTransition(Move{Id: "Min chooses...wisely"}, &GameWithKnownStates{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})
	})
})

type GameWithKnownStates struct {
	isOver bool
	winner Player
	openMoves []minimax.Move
	nextGames []minimax.Game
}

func (game *GameWithKnownStates) AvailableMoves() []minimax.Move {
	return game.openMoves
}

func (game *GameWithKnownStates) Move(move minimax.Move) minimax.Game {
	for i, openMove := range game.openMoves {
		if openMove == move {
			return game.nextGames[i]
		}
	}

	panic("next game not found")
}

func (game *GameWithKnownStates) FindWinner() minimax.Player {
	return game.winner
}

func (game *GameWithKnownStates) IsOver() bool {
	return game.isOver
}

func (game *GameWithKnownStates) AddStateTransition(move Move, nextGame *GameWithKnownStates) {
	game.openMoves = append(game.openMoves, move)
	game.nextGames = append(game.nextGames, nextGame)
}

type Move struct {
	Id string
}

type Player struct {
	Name string
}
