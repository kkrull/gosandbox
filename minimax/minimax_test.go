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
				var spaces = []Space{OpenSpace("A1")}
				var board = NewBoard(spaces, func(Space) string { return "" })
				Expect(board.Minimax()).To(Equal("A1"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces = []Space{ClosedSpace("A1"), OpenSpace("A2")}
				var board = NewBoard(spaces, func(Space) string { return "" })
				Expect(board.Minimax()).To(Equal("A2"))
			})
		})

		Context("when there are 2 or more available spaces with different outcomes", func() {
			It("picks the space that helps the maximizing player win, when it is that player's turn", func() {
				var spaces = []Space{OpenSpace("A1"), OpenSpace("A2")}
				var board = NewBoard(spaces, func(space Space) string {
					switch space.Id() {
					case "A1":
						return "min"
					case "A2":
						return "max"
					}

					return ""
				})
				Expect(board.Minimax()).To(Equal("A2"))
			})
		})
	})
})
