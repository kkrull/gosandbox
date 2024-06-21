package romannumeral_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRomanNumeral(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Romannumeral Suite")
}
