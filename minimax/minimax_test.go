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
	It("returns an outcome", func() {
		game := FakeGame{Over: true}
		Expect(Minimax(game, game.MaximizingPlayer())).To(BeAssignableToTypeOf(Outcome{}))
	})

	Context("given a game that the maximizing player has won", func() {
		It("scores it as +1", func() {
			game := FakeGame{
				Over:   true,
				Winner: maximizer}
			Expect(Minimax(game, game.MaximizingPlayer())).To(BeEquivalentTo(Outcome{BoundedScore: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores it as -1", func() {
			game := FakeGame{
				Over:   true,
				Winner: minimizer}
			Expect(Minimax(game, game.MaximizingPlayer())).To(BeEquivalentTo(Outcome{BoundedScore: -1}))
		})
	})

	Context("given a game that has ended in a draw", func() {
		It("scores it as 0", func() {
			game := FakeGame{Over: true}
			Expect(Minimax(game, game.MaximizingPlayer())).To(BeEquivalentTo(Outcome{BoundedScore: 0}))
		})
	})

	Context("given a game that is not over yet", func() {
		It("picks a move", func() {
			game := FakeGame{Over: false}

			anyPlay := FakePlay{Id: "any play"}
			drawGame := FakeGame{Over: true}
			game.AddAvailableChoice(anyPlay, drawGame)

			Expect(Minimax(game, game.MaximizingPlayer())).To(BeEquivalentTo(Outcome{Play: anyPlay}))
		})
	})

	Context("given a game that can be won with 1 more move", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks the play that makes the maximizing player win", func() {
				game := FakeGame{Over: false}

				playToDraw := FakePlay{Id: "draw"}
				drawGame := FakeGame{Over: true}
				game.AddAvailableChoice(playToDraw, drawGame)

				playToMaxWins := FakePlay{Id: "maxWins"}
				maxWins := FakeGame{
					Over:   true,
					Winner: maximizer}
				game.AddAvailableChoice(playToMaxWins, maxWins)
				Expect(Minimax(game, game.MaximizingPlayer())).To(BeEquivalentTo(Outcome{Play: playToMaxWins, BoundedScore: 1}))
			})
		})

		Context("when it is the minimizing player's turn", func() {
			It("picks the play that makes the maximizing player lose", func() {
				game := FakeGame{Over: false}

				playToDraw := FakePlay{Id: "draw"}
				drawGame := FakeGame{Over: true}
				game.AddAvailableChoice(playToDraw, drawGame)

				playToMaxLoses := FakePlay{Id: "maxLoses"}
				maxLoses := FakeGame{
					Over:   true,
					Winner: minimizer}
				game.AddAvailableChoice(playToMaxLoses, maxLoses)
				Expect(Minimax(game, game.MinimizingPlayer())).
					To(BeEquivalentTo(Outcome{Play: playToMaxLoses, BoundedScore: -1}))
			})
		})
	})

	Context("given a game where a player has to make 2 or more plays to win", func() {
		var (
			rootGame FakeGame
			playToMaxWinning Play
			playToMinWinning Play
		)

		BeforeEach(func() {
			rootGame = FakeGame{Over: false}

			playToMaxWinning = FakePlay{Id: "MaxWinning"}
			maxWinning := FakeGame{Over: false}
			maxWinning.AddAvailableChoice(
				FakePlay{Id: "MaxChokes"},
				FakeGame{Over: true})
			maxWinning.AddAvailableChoice(
				FakePlay{Id: "MaxWins"},
				FakeGame{Over: false, Winner: maximizer})

			playToMinWinning = FakePlay{Id: "MinWinning"}
			minWinning := FakeGame{Over: false}
			minWinning.AddAvailableChoice(
				FakePlay{Id: "MinChokes"},
				FakeGame{Over: true})
			minWinning.AddAvailableChoice(
				FakePlay{Id: "MinWins"},
				FakeGame{Over: false, Winner: minimizer})

			rootGame.AddAvailableChoice(playToMaxWinning, maxWinning)
			rootGame.AddAvailableChoice(playToMinWinning, minWinning)
		})

		It("evaluates players in turn, starting from the maximizing player", func() {
			Expect(Minimax(rootGame, rootGame.MaximizingPlayer())).
				To(BeEquivalentTo(Outcome{Play: playToMaxWinning}))
		})

		It("evaluates players in turn, starting from the minimizing player", func() {
			Expect(Minimax(rootGame, rootGame.MinimizingPlayer())).
				To(BeEquivalentTo(Outcome{Play: playToMinWinning}))
		})
	})
})

type FakeGame struct {
	Choices []Choice
	Over    bool
	Winner  Player
}

func (game *FakeGame) AddAvailableChoice(play Play, resultingGame Game) {
	if game.Choices == nil {
		game.Choices = make([]Choice, 0)
	}

	game.Choices = append(game.Choices, Choice{Play: play, ResultingGame: resultingGame})
}

func (game FakeGame) AvailableChoices() []Choice {
	outcomes := make([]Choice, len(game.Choices))
	copy(outcomes, game.Choices)
	return outcomes
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

type FakePlay struct {
	Id string
}
