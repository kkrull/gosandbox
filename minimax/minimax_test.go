package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Minimax", func() {
	It("returns a result", func() {
		game := FakeGame{}
		Expect(Minimax(game)).To(BeAssignableToTypeOf(Result{}))
	})
})

type FakeGame struct {

}
