package rootCmd_test

import (
	"bytes"

	"github.com/kkrull/cobra-hello/rootCmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rootCmd", func() {
	Describe("Execute", func() {
		It("returns no error, given valid args", func() {
			buf := new(bytes.Buffer)
			err := rootCmd.Execute([]string{}, buf, buf)
			Expect(err).To(BeNil())
		})

		It("does not write to stderr, given valid args", func() {
			errBuf := new(bytes.Buffer)
			outBuf := new(bytes.Buffer)
			rootCmd.Execute([]string{}, errBuf, outBuf)
			Expect(errBuf.String()).To(BeEmpty())
		})
	})
})
