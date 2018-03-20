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
				Expect(board.Minimax("anybody")).To(Equal("open"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces = []Space{ClosedSpace("closed"), OpenSpace("open")}
				var board = NewBoard(spaces, nobodyClaimsVictory)
				Expect(board.Minimax("anybody")).To(Equal("open"))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			It("picks the space where the maximizing player wins, when it is that player's turn", func() {
				var spaces = []Space{OpenSpace("worthless"), OpenSpace("victory")}
				var board = NewBoard(spaces, automaticWin("victory"))
				Expect(board.Minimax("max")).To(Equal("victory"))
			})
		})
	})
})
