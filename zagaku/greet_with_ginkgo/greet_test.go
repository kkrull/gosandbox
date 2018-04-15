package greet_with_ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kkrull/gosandbox/zagaku/greet_with_ginkgo"
)

var _ = Describe("Greet", func() {
  It("Greets the world when no names are given", func() {
    Expect(Greet()).To(Equal("Hello World!"))
  })
})
