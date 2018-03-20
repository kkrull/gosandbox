package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Board", func() {
	It("can be declared with spaces", func() {
		var openSpaces []OpenSpace
		openSpaces = append(openSpaces, OpenSpace{IdValue: "A1"})

		var allSpaces []Space
		for _, space := range openSpaces {
			allSpaces = append(allSpaces, space)
		}

		var board = Board{Spaces: allSpaces}
		Expect(board.Spaces).NotTo(BeNil())
	})

	Describe("Minimax", func() {
		Context("when there is only 1 space", func() {
			It("picks that space", func() {
				var openSpaces []OpenSpace
				openSpaces = append(openSpaces, OpenSpace{IdValue: "A1"})

				var allSpaces []Space
				for _, space := range openSpaces {
					allSpaces = append(allSpaces, space)
				}

				var board = Board{Spaces: allSpaces}
				Expect(board.Minimax()).To(Equal("A1"))
			})
		})

		Context("when there is only 1 available space", func() {
			It("picks that space", func() {
				var spaces []Space
				spaces = append(spaces, ClosedSpace{IdValue: "A1"})
				spaces = append(spaces, OpenSpace{IdValue: "A2"})

				var board = Board{Spaces: spaces}
				Expect(board.Minimax()).To(Equal("A2"))
			})
		})
	})
})

type ClosedSpace struct {
	IdValue string
}

func (ClosedSpace) Id() string {
	panic("implement me")
}

func (ClosedSpace) IsAvailable() bool {
	return false
}

type OpenSpace struct {
	IdValue string
}

func (space OpenSpace) Id() string {
	return space.IdValue
}

func (OpenSpace) IsAvailable() bool {
	return true
}
