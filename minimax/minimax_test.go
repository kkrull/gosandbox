package minimax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/minimax"
)

var _ = Describe("Negamax", func() {
	It("scores a game state", func() {
		score := Negamax()
		Expect(score).To(BeNumerically(">=", -1))
		Expect(score).To(BeNumerically("<=", 1))
	})
})
