package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	It("imports", func() {
		Expect(42).To(Equal(42))
		Minimax()
	})

	It("returns a result", func() {
		Expect(Minimax()).To(BeAssignableToTypeOf(Result{}))
	})
})
