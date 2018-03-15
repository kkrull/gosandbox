package greet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/8thlight/greet"
)

var _ = Describe("Greet", func() {
  It("Greets the world when no names are given", func() {
    Expect(Greet()).To(Equal("Hello World!"))
  })

  It("Greets a single person by their name", func() {
    Expect(Greet("George")).To(Equal("Hello George!"))
  })

  It("Greets two people by name, joining words with 'and'", func() {
    Expect(Greet("George", "Judy")).To(Equal("Hello George and Judy!"))
  })

  It("Greets 3 or more people by a comma-separated list of their names", func() {
    Expect(Greet("George", "Judy", "Astro")).To(Equal("Hello George, Judy, and Astro!"))
  })
})
