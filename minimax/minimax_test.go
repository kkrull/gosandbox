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
})

func CompletedGame() Game {
	return GameState{}
}

type GameState struct {
}

func (GameState) IsOver() bool {
	return true
}
