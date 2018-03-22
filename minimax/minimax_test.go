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

	Context("when there is only 1 available move", func() {
		It("picks that move", func() {
			onlySpace := "Mountain"
			game := OneMoveLeftGame(onlySpace)
			move, _ := Minimax(game)
			Expect(move).To(Equal(onlySpace))
		})
	})
})

func CompletedGame() Game {
	return completedGame{}
}

type completedGame struct{}

func (completedGame) AvailableSpaces() []string {
	return []string{}
}

func (completedGame) IsOver() bool {
	return true
}

func OneMoveLeftGame(space string) Game {
	return incompleteGame{OpenSpaces: []string{space}}
}

type incompleteGame struct {
	OpenSpaces []string
}

func (g incompleteGame) AvailableSpaces() []string {
	return g.OpenSpaces
}

func (incompleteGame) IsOver() bool {
	return false
}
