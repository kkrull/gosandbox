package rootCmd_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rootCmd", func() {
	Describe("Execute", func() {
		It("exists", func() {
			Expect(1).To(Equal(1))
		})
	})
})
