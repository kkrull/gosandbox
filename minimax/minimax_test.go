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
		var (
			board Board
			space string
			score int
		)

		Context("when there is only 1 space", func() {
			It("picks that space", func() {
				spaces := []Space{OpenSpace("open")}
				board = NewBoard("max", "min", spaces, nobodyClaimsVictory)
				space, _ = board.Minimax("anybody")
				Expect(space).To(Equal("open"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				spaces := []Space{ClosedSpace("closed"), OpenSpace("open")}
				board = NewBoard("max", "min", spaces, nobodyClaimsVictory)
				space, _ = board.Minimax("anybody")
				Expect(space).To(Equal("open"))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			BeforeEach(func() {
				spaces := []Space{OpenSpace("worthless"), OpenSpace("victory")}
				board = NewBoard("max", "min", spaces, automaticWin("victory"))
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

			Context("when it is the minimizing player's turn", func() {
				It("picks the space where the maximizing player loses", func() {
					space, score = board.Minimax("min")
					Expect(space).To(Equal("victory"))
				})
				It("returns a negative score for that space", func() {
					space, score = board.Minimax("min")
					Expect(score).To(Equal(-1))
				})
			})
		})
	})
})
