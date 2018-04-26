package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max    = Player{Name: "Max"}
	min    = Player{Name: "Min"}
	scorer = Minimax{
		MaximizingPlayer: max,
		MinimizingPlayer: min,
	}
)

var _ = Describe("Minimax", func() {
	Describe("#Score", func() {
		It("scores a game ending in a draw as 0", func() {
			game := &FakeGame{isOver: true}
			Expect(scorer.Score(game, max)).To(Equal(0))
		})

		It("scores a game won by the maximizing player as +1", func() {
			game := &FakeGame{isOver: true, winner: max}
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("scores a game won by the minimizing player as -1", func() {
			game := &FakeGame{isOver: true, winner: min}
			Expect(scorer.Score(game, max)).To(Equal(-1))
		})

		It("maximizes the score for the maximizing player, among all possible outcomes", func() {
			game := &FakeGame{}
			game.AddNextGame(Move{Id: "Max Chooses....Poorly"}, &FakeGame{isOver: true})
			game.AddNextGame(Move{Id: "Max Chooses....Wisely"}, &FakeGame{isOver: true, winner: max})
			Expect(scorer.Score(game, max)).To(Equal(1))
		})

		It("minimizes the score for the minimizing player, among all possible outcomes", func() {
			game := &FakeGame{}
			game.AddNextGame(Move{Id: "Min Chooses....Poorly"}, &FakeGame{isOver: true})
			game.AddNextGame(Move{Id: "Min Chooses....Wisely"}, &FakeGame{isOver: true, winner: min})
			Expect(scorer.Score(game, min)).To(Equal(-1))
		})
	})
})

type FakeGame struct {
	isOver    bool
	winner    Player
	nextMoves []Move
	nextGames []Game
}

func (game *FakeGame) AvailableMoves() []Move {
	return game.nextMoves
}

func (game *FakeGame) Move(selectedMove Move) Game {
	for i, move := range game.nextMoves {
		if move == selectedMove	{
			return game.nextGames[i]
		}
	}

	panic("Unrecognized move")
}

func (game *FakeGame) FindWinner() Player {
	return game.winner
}

func (game *FakeGame) IsOver() bool {
	return game.isOver
}

func (game *FakeGame) AddNextGame(move Move, nextGame *FakeGame) {
	game.nextMoves = append(game.nextMoves, move)
	game.nextGames = append(game.nextGames, nextGame)
}
