package primitive_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// START OMIT
var _ = Describe("Channel", func() {
	var (
		wrongAnswer = 41
		rightAnswer = 42
	)

	It("acts as a buffered FIFO queue", func() {
		answers := make(chan int, 2) // The addition of a capacity makes it a buffered channel // HL
		answers <- wrongAnswer
		answers <- rightAnswer

		Expect(<-answers).To(Equal(wrongAnswer)) //41
		Expect(<-answers).To(Equal(rightAnswer)) //42
	})

	It("blocks when unbuffered", func(done Done) { //Note the "Done" object (like Mocha) // HL
		answers := make(chan string)
		answers <- "one at a time"
		answers <- "will block until somebody receives the first answer" // Blocks // HL
		close(done) // Signal that the test is over by closing the Done channel // HL
					// Except it never gets here...
	})
})
// END OMIT