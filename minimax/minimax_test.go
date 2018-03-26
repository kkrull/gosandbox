package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

func GameOver() Game {
	return gameOver{}
}

type gameOver struct { }

func (gameOver) IsOver() bool {
	return true
}

var _ = Describe("Minimax", func() {
	It("exists", func() {
		Minimax(GameOver())
	})

	Context("when the game is over", func() {
		It("returns an error", func() {
			game := GameOver()
			_, err := Minimax(game)
			Expect(err).To(MatchError("the game is already over"))
		})
	})
})
