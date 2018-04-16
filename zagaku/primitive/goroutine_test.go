package primitive_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// START OMIT
var _ = Describe("goroutines", func() {
	It("runs a function asynchronously", func(done Done) {
		buffer := make(chan string, 1)
		go func() { // Just say the word "go" and this function runs asynchronously // HL
			time.Sleep(1 * time.Second)
			buffer <- "result"
		}()

		select { //Make sure this code runs before the goroutine finishes
		case immediatelyAvailableValue, valueReceived := <-buffer: //Non-blocking receive
			//We would get here if the goroutine executed sequentially
			Expect(valueReceived).To(BeFalse()) //...and this would fail the test
			Expect(immediatelyAvailableValue).To(BeEmpty())
		default:
		}

		computedResult := <-buffer // Blocks until the result is sent by the goroutine // HL
		Expect(computedResult).To(Equal("result"))
		close(done)
	}, 2)
})
// END OMIT