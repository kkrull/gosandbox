package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Board", func() {
	It("can be declared with spaces", func() {
		var spaces []Space
		var board = Board{Spaces: spaces}
		Expect(board).NotTo(BeNil())
	})
})
