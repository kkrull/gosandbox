package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("when the game is over", func() {
		It("returns an error", func() {
			game := CompletedGame()
			_, err := Minimax(game)
			Expect(err).To(MatchError("minimax: game over"))
		})
	})

	Context("when there is only 1 available space", func() {
		It("picks that space", func() {
			game := KingOfTheMountainGame("Mountain")
			Expect(Minimax(game)).To(Equal("Mountain"))
		})
	})

	Context("when there are 2 of more available spaces", func() {
		Context("when it is the maximizing player's turn", func() {
			It("picks the move that causes the maximizing player to win", func() {
				game := CaptureTheFlagGame(EmptySpace("_"), FlagSpace("F"))
				Expect(Minimax(game)).To(Equal("F"))
			})
		})
	})
})
