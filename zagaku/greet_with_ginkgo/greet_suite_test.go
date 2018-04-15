package greet_with_ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGreet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greet Suite")
}
