package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

func nobodyClaimsVictory(string, Space) string { return "" }
func automaticWin(id string) func(string, Space) string {
	return func(ifPlayer string, claimsSpace Space) string {
		if claimsSpace.Id() == id {
			return ifPlayer
		} else {
			return ""
		}
	}
}

var _ = Describe("Board", func() {
	Describe("#Minimax", func() {
		Context("when there is only 1 space", func() {
			It("picks that space", func() {
				var spaces = []Space{OpenSpace("open")}
				var board = NewBoard(spaces, nobodyClaimsVictory)
				var space, _ = board.Minimax("anybody")
				Expect(space).To(Equal("open"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces = []Space{ClosedSpace("closed"), OpenSpace("open")}
				var board = NewBoard(spaces, nobodyClaimsVictory)
				var space, _ = board.Minimax("anybody")
				Expect(space).To(Equal("open"))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			var (
				board Board
				space string
				score int
			)

			BeforeEach(func() {
				var spaces = []Space{OpenSpace("worthless"), OpenSpace("victory")}
				board = NewBoard(spaces, automaticWin("victory"))
			})

			Context("when it is the maximizing player's turn", func() {
				It("picks the space where the maximizing player wins", func() {
					space, score = board.Minimax("max")
					Expect(space).To(Equal("victory"))
				})
				It("returns a positive score for that space", func() {
					space, score = board.Minimax("max")
					Expect(score).To(Equal(1))
				})
			})

			//It("picks the space where the minimizing player wins, when it is that player's turn", func() {
			//	var spaces = []Space{OpenSpace("worthless"), OpenSpace("victory")}
			//	var board = NewBoard(spaces, automaticWin("victory"))
			//	var space, score = board.Minimax("min")
			//	Expect(space).To(Equal("victory"))
			//	Expect(score).To(Equal(-1))
			//})
		})
	})
})
