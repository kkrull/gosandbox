package main

import (
	"fmt"
	"os"

	"github.com/kkrull/cobra-hello/rootCmd"
)

func main() {
	fmt.Println("Raw args:")
	for i, rawArg := range os.Args {
		fmt.Printf("- %d: %s\n", i, rawArg)
	}

	err := rootCmd.Execute(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
