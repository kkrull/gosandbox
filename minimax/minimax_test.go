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

	Context("when there are multiple moves to consider", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks the space that leads to immediate victory for the maximizing player", func() {
				game := newMultiSpaceGame()
				Expect(Minimax(game, "Max")).To(BeEquivalentTo(Result{
					Space: "Victory",
					Score: 1,
				}))
			})
		})
	})

	XIt("Switches players")
})

func newMultiSpaceGame() multiSpaceGame {
	return multiSpaceGame{
		spaces: []string{"Useless", "Victory"},
		owners: make([]string, 2),
	}
}

type multiSpaceGame struct {
	spaces []string
	owners []string
}

func (game multiSpaceGame) HasWinner() bool {
	return game.FindWinner() != ""
}

func (game multiSpaceGame) FindWinner() string {
	return game.owners[1]
}

func (game multiSpaceGame) IsOver() bool {
	if game.HasWinner() {
		return true
	}

	return len(game.AvailableSpaces()) == 0
}

func (game multiSpaceGame) AvailableSpaces() []string {
	openSpaces := make([]string, 0)
	for i, owner := range game.owners {
		if owner == "" {
			openSpaces = append(openSpaces, game.spaces[i])
		}
	}

	return openSpaces
}

func (game multiSpaceGame) LegalMoves(player string) []Move {
	if game.owners[1] != "" {
		return []Move{}
	}

	open := make([]Move, 0)
	for i, owner := range game.owners {
		if owner == "" {
			nextOwners := make([]string, len(game.owners))
			copy(nextOwners, game.owners)
			nextOwners[i] = player

			move := Move{
				ClaimedSpace: game.spaces[i],
				Game: multiSpaceGame{
					spaces: game.spaces,
					owners: nextOwners}}
			open = append(open, move)
		}
	}

	return open
}
