package primitive_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestZeroValue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Primitive Suite")
}
