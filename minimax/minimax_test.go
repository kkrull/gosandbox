package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("given a game that the maximizing player has won", func() {
		It("scores that game as +1", func() {
			game := CompletedGame{Winner: "Max"}
			Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{Score: 1}))
		})
	})

	Context("given a game that the minimizing player has won", func() {
		It("scores that game as -1", func() {
			game := CompletedGame{Winner: "Min"}
			Expect(Minimax(game, "Min")).To(BeEquivalentTo(Result{Score: -1}))
		})
	})

	Context("given a game ending in a draw", func() {
		It("scores that game as -1", func() {
			game := CompletedGame{}
			Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{Score: 0}))
		})
	})

	Context("given a game that is not over yet", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks a move where the maximizing player wins", func() {
				game := DuelGame{Space: "Victory"}
				Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{
					Space: "Victory",
					Score: 1,
				}))
			})
		})

		Context("when it is the minimizing player's turn", func() {
			It("picks a move where the maximizing player loses", func() {
				game := DuelGame{Space: "Victory"}
				Expect(Minimax(game, "Min")).To(BeEquivalentTo(Result{
					Space: "Victory",
					Score: -1,
				}))
			})
		})
	})
})

type CompletedGame struct {
	Winner string
}

func (game CompletedGame) FindWinner() string {
	return game.Winner
}

func (CompletedGame) IsOver() bool {
	return true
}

func (game CompletedGame) LegalMoves(player string) []Move {
	return []Move{}
}

type DuelGame struct {
	Space string
}

func (DuelGame) FindWinner() string {
	return ""
}

func (DuelGame) IsOver() bool {
	return false
}

func (game DuelGame) LegalMoves(player string) []Move {
	return []Move{
		Move{
			Game: CompletedGame{Winner: player},
			ClaimedSpace: game.Space,
		},
	}
}
