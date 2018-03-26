package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("when the game is over", func() {
		It("returns an error", func() {
			game := GameOver()
			_, err := Minimax(game)
			Expect(err).To(MatchError("the game is already over"))
		})
	})

	Context("when the game is still going", func() {
		It("does not return an error", func() {
			game := QuickDrawGame("Victory")
			_, err := Minimax(game)
			Expect(err).To(BeNil())
		})

		It("returns an available space", func() {
			game := QuickDrawGame("Victory")
			result, _ := Minimax(game)
			Expect(result).To(Equal("Victory"))
		})
	})
})
