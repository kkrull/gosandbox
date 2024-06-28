package rootCmd_test

import (
	"bytes"

	"github.com/kkrull/cobra-hello/rootCmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	errBuffer = new(bytes.Buffer)
	outBuffer = new(bytes.Buffer)
)

var _ = Describe("rootCmd", func() {
	Describe("Execute", func() {
		It("returns no error, given valid args", func() {
			err := rootCmd.Execute([]string{}, errBuffer, outBuffer)
			Expect(err).To(BeNil())
		})

		It("does not write to stderr, given valid args", func() {
			rootCmd.Execute([]string{}, errBuffer, outBuffer)
			Expect(errBuffer.String()).To(BeEmpty())
		})
	})
})
