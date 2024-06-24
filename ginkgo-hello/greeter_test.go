package ginkgohello_test

import (
	ginkgohello "github.com/kkrull/ginkgohello"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var subject *ginkgohello.Greeter

var _ = Describe("Greeter", func() {
	Describe("MakeGreeting", func() {
		It("greets the world", func() {
			subject = &ginkgohello.Greeter{}
			Expect(subject.Greet()).To(Equal("Hello World!"))
		})

		It("greets a person by name", func() {
			subject = &ginkgohello.Greeter{"George"}
			Expect(subject.Greet()).To(Equal("Hello George!"))
		})
	})
})
