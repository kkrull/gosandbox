package romannumeral_test

import (
	. "github.com/kkrull/gosandbox/romannumeral"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RomanNumeral", func() {
	Describe("#Convert", func() {
		It("converts 1 to I", func() {
			Expect(RomanNumeral(1)).To(Equal("I"))
		})
	})
})
