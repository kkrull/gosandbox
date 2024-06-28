package main_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("cobra.Command", func() {
	Describe("#Execute", func() {
		It("parses flags", func() {
			subject := &cobra.Command{}
			buf := new(bytes.Buffer)
			subject.SetErr(buf)
			subject.SetOut(buf)

			optionalStringFlag := subject.Flags().String("optional", "", "")

			subject.SetArgs([]string{"--optional", "some-value"})
			err := subject.Execute()
			Expect(err).To(BeNil())

			Expect(*optionalStringFlag).To(Equal("some-value"))
		})
	})
})
