package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	quit := make(chan bool, 1)
	go func() {
		for sig := range c {
			fmt.Printf("Received signal: %s\n", sig.String())
			quit <- true
		}
	}()

	fmt.Println("Press Ctrl+C to exit")
	<-quit
	fmt.Println("Goodbye")
}
