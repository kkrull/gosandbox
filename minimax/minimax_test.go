package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	Context("when the game is over", func() {
		It("returns an error", func() {
			_, err := Minimax()
			Expect(err).To(MatchError("minimax: game over"))
		})
	})
})
