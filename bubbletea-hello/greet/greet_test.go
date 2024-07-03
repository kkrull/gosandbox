package greet_test

import (
	"github.com/kkrull/bubbletea-hello/greet"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("greeterModel", func() {
	Describe("::NewModel", func() {
		It("returns a model", func() {
			subject := greet.NewModel()
			Expect(subject).NotTo(BeNil())
		})
	})
})
