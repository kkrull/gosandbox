package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	maximizer = Player{Name: "Max"}
	minimizer = Player{Name: "Min"}
)

var _ = Describe("Minimax", func() {
	It("returns a result", func() {
		game := FakeGame{Over: true}
		Expect(Minimax(game)).To(BeAssignableToTypeOf(Result{}))
	})

	Context("given a game that the maximizing player has won", func() {
		It("scores it as +1", func() {
			game := FakeGame{
				Over:   true,
				Winner: maximizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores it as -1", func() {
			game := FakeGame{
				Over:   true,
				Winner: minimizer}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: -1}))
		})
	})

	Context("given a game that has ended in a draw", func() {
		It("scores it as 0", func() {
			game := FakeGame{Over: true}
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Score: 0}))
		})
	})

	Context("given a game that is not over yet", func() {
		It("picks a move", func() {
			game := FakeGame{Over: false}

			anyMove := FakeMove{Id: "any move"}
			drawGame := FakeGame{Over: true}
			game.AddAvailableMove(anyMove, drawGame)

			Expect(Minimax(game)).To(BeEquivalentTo(Result{Move: anyMove}))
		})
	})

	Context("given a game that can be won with 1 more move", func() {
		It("picks the move that makes the maximizing player win", func() {
			game := FakeGame{Over: false}

			moveToDraw := FakeMove{Id: "draw"}
			drawGame := FakeGame{Over: true}
			game.AddAvailableMove(moveToDraw, drawGame)

			moveToMaxWins := FakeMove{Id: "maxWins"}
			maxWins := FakeGame{
				Over:   true,
				Winner: maximizer}
			game.AddAvailableMove(moveToMaxWins, maxWins)
			Expect(Minimax(game)).To(BeEquivalentTo(Result{Move: moveToMaxWins, Score: 1}))
		})
	})
})

type FakeGame struct {
	Moves  []Outcome
	Over   bool
	Winner Player
}

type Outcome struct {
	Move     Move
	NextGame Game
}

func (game *FakeGame) AddAvailableMove(move Move, nextGame Game) {
	if game.Moves == nil {
		game.Moves = make([]Outcome, 0)
	}

	game.Moves = append(game.Moves, Outcome{Move: move, NextGame: nextGame})
}

func (game FakeGame) AvailableMoves() []Move {
	moves := make([]Move, len(game.Moves))
	for i, move := range game.Moves {
		moves[i] = move.Move
	}

	return moves
}

func (game FakeGame) FindWinner() Player {
	return game.Winner
}

func (game FakeGame) IsOver() bool {
	return game.Over
}

func (game FakeGame) MaximizingPlayer() Player {
	return maximizer
}

func (game FakeGame) MinimizingPlayer() Player {
	return minimizer
}

type FakeMove struct {
	Id string
}
