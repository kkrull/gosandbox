package stringcalc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringcalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringcalc Suite")
}
