package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	uselessSpace := "Useless"
	playerWins := "PlayerWins"
	opponentWins := "OpponentWins"

	Context("when the game is already over", func() {
		It("returns an error", func() {
			game := EndGame()
			_, err := Minimax(game, "Player")
			Expect(err).To(MatchError("game is already over"))
		})
	})

	Context("when the game is not over yet", func() {
		It("does not return an error", func() {
			game := OpenGame(42)
			_, err := Minimax(game, "Player")
			Expect(err).NotTo(HaveOccurred())
		})

		It("picks a space in the game", func() {
			game := OpenGame(42)
			space, _ := Minimax(game, "Player")
			Expect(space).NotTo(BeZero())
		})

		It("picks an available space in the game", func() {
			openSpace := "Available"
			game := OpenGame(openSpace)
			space, _ := Minimax(game, "Player")
			Expect(space).To(Equal(openSpace))
		})
	})

	Context("when there is a space that leads to an instant victory", func() {
		It("picks the space that wins the game", func() {
			game := OpenGame(uselessSpace, playerWins)
			space, _ := Minimax(game, "Player")
			Expect(space).To(Equal(playerWins))
		})
	})

	Context("when each player benefits from picking a different space", func() {
		It("picks the space that benefits that player", func() {
			game := OpenGame(playerWins, opponentWins)
			space, _ := Minimax(game, "Player")
			Expect(space).To(Equal(playerWins))
		})
	})

	Context("when it is the opponent's turn", func() {
		It("picks the space that makes the maximizing player lose", func() {
			game := OpenGame(playerWins, opponentWins)
			space, _ := Minimax(game, "Opponent")
			Expect(space).To(Equal(opponentWins))
		})
	})

	Context("when each player needs to pick 2 or more spaces to win", func() {
		XIt("picks the first of those two spaces", func() {
			game := TwoStepGame("PlayerStep1", uselessSpace, uselessSpace)
			space, _ := Minimax(game, "Player")
			Expect(space).To(Equal("PlayerStep1"))
		})
	})
})
