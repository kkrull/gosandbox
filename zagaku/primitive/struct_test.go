package primitive_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// START OMIT
type Animal struct {
	Id int
	Name string
	Type string
}

var _ = Describe("structs", func() {
	It("have a zero value in which each field is zero-valued", func() {
		zeroAnimal := Animal{}
		Expect(zeroAnimal).NotTo(BeNil())
		Expect(zeroAnimal.Id).To(Equal(0))
		Expect(zeroAnimal.Name).To(BeEmpty())
	})

	It("can be initialized with 1 or more fields", func() {
		anonymousAnimal := Animal{Type: "mammal"}
		Expect(anonymousAnimal.Name).To(BeEmpty())
		Expect(anonymousAnimal.Type).To(Equal("mammal"))
	})
})
// END OMIT