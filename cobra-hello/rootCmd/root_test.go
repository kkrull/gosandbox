package rootCmd_test

import (
	"github.com/kkrull/cobra-hello/rootCmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rootCmd", func() {
	Describe("Execute", func() {
		It("returns no error, given valid args", func() {
			err := rootCmd.Execute([]string{})
			Expect(err).To(BeNil())
		})
	})
})
