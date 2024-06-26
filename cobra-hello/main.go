package main

import (
	"fmt"
	"os"

	"github.com/kkrull/cobra-hello/rootCmd"
)

func main() {
	fmt.Println("Raw args:")
	for i, v := range os.Args {
		fmt.Printf("- %d: %s\n", i, v)
	}

	rootCmd.Execute()
}
