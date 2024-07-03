package greet_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGreet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "greet")
}
