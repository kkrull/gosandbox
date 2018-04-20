package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max = Player{Name: "max"}
	min = Player{Name: "min"}
)

var _ = Describe("Minimax", func() {
	It("scores a game ending in a draw as 0", func() {
		game := &FakeGame{isOver: true}
		Expect(Minimax(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as +1", func() {
		game := &FakeGame{isOver: true, winner: max}
		Expect(Minimax(game, max)).To(Equal(1))
	})

	It("scores a game won by the minimizing player as -1", func() {
		game := &FakeGame{isOver: true, winner: min}
		Expect(Minimax(game, max)).To(Equal(-1))
	})

	It("returns the maximum score that the maximizing player can achieve in a non-terminal game", func() {
		game := &FakeGame{nextPossibleGames: []Game{
			&FakeGame{isOver: true},
			&FakeGame{isOver: true, winner: max},
		}}
		Expect(Minimax(game, max)).To(Equal(1))
	})

	It("returns the minimum score that the minimizing player can achieve in a non-terminal game", func() {
		game := &FakeGame{nextPossibleGames: []Game{
			&FakeGame{isOver: true},
			&FakeGame{isOver: true, winner: min},
		}}
		Expect(Minimax(game, min)).To(Equal(-1))
	})
})

type FakeGame struct {
	isOver bool
	winner Player
	nextPossibleGames []Game
}

func (game *FakeGame) FindWinner() Player {
	return game.winner
}

func (game *FakeGame) MaximizingPlayer() Player {
	return max
}

func (game *FakeGame) MinimizingPlayer() Player {
	return min
}

func (game *FakeGame) NextPossibleGames() []Game {
	return game.nextPossibleGames
}

func (game *FakeGame) IsOver() bool {
	return game.isOver
}
