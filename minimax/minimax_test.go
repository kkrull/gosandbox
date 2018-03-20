package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Board", func() {
	It("can be declared with spaces", func() {
		var spaces = []Space { OpenSpace("A1") }
		var board = Board{Spaces: spaces}
		Expect(board.Spaces).NotTo(BeNil())
	})

	Describe("Minimax", func() {
		Context("when there is only 1 space", func() {
			It("picks that space", func() {
				var spaces = []Space { OpenSpace("A1") }
				var board = Board{Spaces: spaces}
				Expect(board.Minimax()).To(Equal("A1"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces = []Space { ClosedSpace("A1"), OpenSpace("A2") }
				var board = Board{Spaces: spaces}
				Expect(board.Minimax()).To(Equal("A2"))
			})
		})

		Context("when there are 2 or more spaces with different outcomes", func() {
			//It("picks the space that helps the maximizing player win", func() {
			//	var spaces = []Space {
			//		OpenSpace{IdValue: "A1", WinningPlayer: "min"},
			//		OpenSpace{IdValue: "A2", WinningPlayer: "max"},
			//	}
			//
			//	var board = Board{Spaces: spaces}
			//	Expect(board.Minimax()).To(Equal("A2"))
			//})
		})
	})
})
