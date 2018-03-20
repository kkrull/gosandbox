package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Board", func() {
	Describe("#Minimax", func() {
		Context("when there is only 1 space", func() {
			It("picks that space", func() {
				var spaces = []Space{OpenSpace("open")}
				var board = NewBoard(spaces, func(string, Space) string { return "" })
				Expect(board.Minimax("anybody")).To(Equal("open"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces = []Space{ClosedSpace("closed"), OpenSpace("open")}
				var board = NewBoard(spaces, func(string, Space) string { return "" })
				Expect(board.Minimax("anybody")).To(Equal("open"))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			It("picks the space where the maximizing player wins, when it is that player's turn", func() {
				var spaces = []Space{OpenSpace("worthless"), OpenSpace("victory")}
				var board = NewBoard(spaces, func(player string, space Space) string {
					if space.Id() == "victory" {
						return player
					} else {
						return ""
					}
				})

				Expect(board.Minimax("max")).To(Equal("victory"))
			})
		})
	})
})
