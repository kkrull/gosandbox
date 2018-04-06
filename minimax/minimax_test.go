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
	It("returns an Outcome", func() {
		game := FakeGame{Winner: max}
		Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Score: 1}))
	})

	It("scores a game the maximizing player has won as +1", func() {
		game := FakeGame{Winner: max}
		Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Score: 1}))
	})

	It("scores a game the maximizing player has won as -1", func() {
		game := FakeGame{Winner: min}
		Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Score: -1}))
	})

	It("scores a game ending in a draw as 0", func() {
		game := FakeGame{Over: true}
		Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Score: 0}))
	})

	Context("given a game that is not over yet", func() {
		It("picks an available play", func() {
			game := FakeGame{}
			anyPlay := Play{Id: "any"}
			game.AddAvailablePlay(anyPlay, FakeGame{})
			Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Play: anyPlay}))
		})
	})

	Context("given a game that can be won by the maximizing player", func() {
		It("the maximizing player picks the play that helps it win the game", func() {
			game := FakeGame{}
			drawPlay := Play{Id: "draw"}
			maxWinsPlay := Play{Id: "MaxWins"}

			game.AddAvailablePlay(drawPlay, FakeGame{Over: true})
			game.AddAvailablePlay(maxWinsPlay, FakeGame{Winner: max})

			Expect(Minimax(game)).To(BeEquivalentTo(Outcome{Play: maxWinsPlay, Score: 1}))
		})
	})
})

type FakeGame struct {
	Over           bool
	Winner         Player
	availablePlays []Choice
}

func (game FakeGame) IsOver() bool {
	return game.Over
}

func (game *FakeGame) AddAvailablePlay(play Play, nextGame Game) {
	game.availablePlays = append(game.availablePlays, Choice{
		Play:     play,
		NextGame: nextGame,
	})
}

func (game FakeGame) AvailablePlays() []Choice {
	return game.availablePlays
}

func (game FakeGame) FindWinner() Player {
	return game.Winner
}

func (game FakeGame) MaximizingPlayer() Player {
	return max
}

func (game FakeGame) MinimizingPlayer() Player {
	return min
}
