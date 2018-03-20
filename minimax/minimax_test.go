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
			BeforeEach(func() {
				board = NewBoard("max", "min", []Space{}, nobodyClaimsVictory)
				space, score, err = board.Minimax("max")
			})

			It("returns an error", func() {
				Expect(err).To(MatchError("the game is already over"))
			})
			It("returns zero values for everything else", func() {
				Expect(space).To(Equal(""))
				Expect(score).To(Equal(0))
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
			It("does not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			Context("when it is the maximizing player's turn", func() {
				BeforeEach(func() {
					spaces := []Space{OpenSpace("worthless"), OpenSpace("victory")}
					board = NewBoard("max", "min", spaces, automaticWin("victory"))
					space, score, err = board.Minimax("max")
				})

				It("picks the space where the maximizing player wins", func() {
					Expect(space).To(Equal("victory"))
				})
				It("returns a positive score for that space", func() {
					Expect(score).To(Equal(1))
				})
				It("does not return an error", func() {
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when it is the minimizing player's turn", func() {
				BeforeEach(func() {
					spaces := []Space{OpenSpace("worthless"), OpenSpace("victory")}
					board = NewBoard("max", "min", spaces, automaticWin("victory"))
					space, score, err = board.Minimax("min")
				})

				It("picks the space where the maximizing player loses", func() {
					Expect(space).To(Equal("victory"))
				})
				It("returns a negative score for that space", func() {
					Expect(score).To(Equal(-1))
				})
				It("does not return an error", func() {
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
	})
})
