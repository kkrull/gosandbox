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
			err error
		)

		Context("when there are not any available spaces left", func() {
			It("returns an error, with zero values for everything else", func() {
				board = NewBoard("max", "min", []Space{}, nobodyClaimsVictory)
				space, score, err = board.Minimax("max")
				Expect(space).To(Equal(""))
				Expect(score).To(Equal(0))
				Expect(err).To(MatchError("the game is already over"))
			})
		})

		Context("when none of the available spaces result in victory for either player", func() {
			BeforeEach(func() {
				spaces := []Space{ClosedSpace("closed"), OpenSpace("open")}
				board = NewBoard("max", "min", spaces, nobodyClaimsVictory)
				space, score, err = board.Minimax("max")
			})

			It("returns the first available space", func() {
				Expect(space).To(Equal("open"))
			})
			It("returns 0 for the score", func() {
				Expect(score).To(Equal(0))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			BeforeEach(func() {
				spaces := []Space{OpenSpace("worthless"), OpenSpace("victory")}
				board = NewBoard("max", "min", spaces, automaticWin("victory"))
			})

			Context("when it is the maximizing player's turn", func() {
				It("picks the space where the maximizing player wins", func() {
					space, score, err = board.Minimax("max")
					Expect(space).To(Equal("victory"))
				})
				It("returns a positive score for that space", func() {
					space, score, err = board.Minimax("max")
					Expect(score).To(Equal(1))
				})
			})

			Context("when it is the minimizing player's turn", func() {
				It("picks the space where the maximizing player loses", func() {
					space, score, _ = board.Minimax("min")
					Expect(space).To(Equal("victory"))
				})
				It("returns a negative score for that space", func() {
					space, score, _ = board.Minimax("min")
					Expect(score).To(Equal(-1))
				})
			})
		})
	})
})
