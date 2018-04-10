package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	It("returns a score", func() {
		Expect(Minimax()).To(BeNumerically(">=", -1))
		Expect(Minimax()).To(BeNumerically("<=", 1))
	})
})
