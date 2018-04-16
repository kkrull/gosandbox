package greet_with_ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// START OMIT
//greet_suite_test.go
func TestGreet(t *testing.T) { //Integration with built-in testing // HL
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greet Suite")
}
// END OMIT
