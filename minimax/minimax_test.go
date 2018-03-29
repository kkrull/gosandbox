package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	It("Exists", func() {
		Expect(Minimax()).To(BeZero())
	})
})
