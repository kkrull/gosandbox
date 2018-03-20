package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

type OpenSpace struct {
	IdValue string
}

func (OpenSpace) Id() string {
	panic("implement me")
}

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
})
