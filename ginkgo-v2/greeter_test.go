package ginkgo_v2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	ginkgo_v2 "kkrull.github.com/gosandbox/ginkgov2"
)

var subject *ginkgo_v2.Greeter

var _ = Describe("Greeter", func() {
	Describe("MakeGreeting", func() {
		It("greets the world", func() {
			subject = &ginkgo_v2.Greeter{}
			Expect(subject.Greet()).To(Equal("Hello World!"))
		})

		It("greets a person by name", func() {
			subject = &ginkgo_v2.Greeter{ "George" }
			Expect(subject.Greet()).To(Equal("Hello George!"))
		})
	})
})
