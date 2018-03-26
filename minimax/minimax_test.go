package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	var (
		game  Game
		space Space
		err   error
	)

	Context("when the game is over", func() {
		It("returns an error", func() {
			game = GameOver()
			space, err = Minimax(game)
			Expect(err).To(MatchError("the game is already over"))
		})
	})

	Context("when the game is still going", func() {
		BeforeEach(func() {
			game = QuickDrawGame("Victory")
			space, err = Minimax(game)
		})

		It("returns an available space", func() {
			Expect(space).To(Equal("Victory"))
		})
		It("does not return an error", func() {
			Expect(err).To(BeNil())
		})
	})
})
