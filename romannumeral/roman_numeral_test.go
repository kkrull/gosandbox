package romannumeral_test

import (
	"github.com/kkrull/romannumeral"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RomanNumeral", func() {
	Describe("#Convert", func() {
		It("converts 1 to I", func() {
			Expect(romannumeral.RomanNumeral(1)).To(Equal("I"))
		})
	})
})
