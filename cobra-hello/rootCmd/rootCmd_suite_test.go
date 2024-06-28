package rootCmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRootCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "rootCmd")
}
