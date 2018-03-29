package main

import "fmt"

func main() {
	channel := make(chan int, 1)
	go func() {
		input := 42
		fmt.Printf("Writing to channel: %d\n", input)
		channel <- input
	}()

	output := <-channel
	fmt.Printf("Read from channel: %d\n", output)
}
