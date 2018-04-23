package minimax_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var (
	max    = Player{Name: "Max"}
	min    = Player{Name: "Min"}
	game   Game
	scorer = &Scorer{GameStatesFromNextMove: &FakeGameTransitions{}}
)

var _ = Describe("Score", func() {
	It("scores a game ending in a draw as 0", func() {
		game = &FakeGame{isOver: true}
		Expect(scorer.Score(game, max)).To(Equal(0))
	})

	It("scores a game won by the maximizing player as +1", func() {
		game = &FakeGame{isOver: true, winner: max}
		Expect(scorer.Score(game, max)).To(Equal(1))
	})

	It("scores a game won by the minimizing player as -1", func() {
		game = &FakeGame{isOver: true, winner: min}
		Expect(scorer.Score(game, max)).To(Equal(-1))
	})

	It("returns the maximum possible score for the maximizing player, given an unfinished game", func() {
		game = &FakeGame{
			nextGames: []Game{
				&FakeGame{isOver: true},
				&FakeGame{isOver: true, winner: max},
			},
			nextMoves: []Move{
				Move{Id: "max chooses poorly"},
				Move{Id: "max chooses wisely"},
			},
		}
		Expect(scorer.Score(game, max)).To(Equal(1))
	})

	It("returns the minimum possible score for the minimizing player, given an unfinished game", func() {
		game = &FakeGame{
			nextGames: []Game{
				&FakeGame{isOver: true},
				&FakeGame{isOver: true, winner: min},
			},
			nextMoves: []Move{
				Move{Id: "min chooses poorly"},
				Move{Id: "min chooses wisely"},
			},
		}
		Expect(scorer.Score(game, min)).To(Equal(-1))
	})
})

type FakeGameTransitions struct {

}


func (transition *FakeGameTransitions) GameStatesFromNextMove(game Game) []Game {
	nextGames := make([]Game, 0)
	for _, move := range game.AvailableMoves() {
		nextGame := game.Move(move)
		nextGames = append(nextGames, nextGame)
	}

	return nextGames
}


type FakeGame struct {
	isOver    bool
	winner    Player
	nextGames []Game
	nextMoves []Move
}

func (game *FakeGame) AvailableMoves() []Move {
	return game.nextMoves
}

func (game *FakeGame) Move(move Move) Game {
	for i, nextMove := range game.nextMoves {
		if move == nextMove {
			return game.nextGames[i]
		}
	}

	panic(fmt.Sprintf("next game not found for move %v", move))
}

func (game *FakeGame) FindWinner() Player {
	return game.winner
}

func (game *FakeGame) IsOver() bool {
	return game.isOver
}

func (game *FakeGame) MaximizingPlayer() Player {
	return max
}

func (game *FakeGame) MinimizingPlayer() Player {
	return min
}
