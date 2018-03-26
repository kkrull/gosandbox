package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	fmt.Printf("Arguments: %v\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("- %d: %s\n", i, arg)
	}

	flagSet := flag.NewFlagSet("custom program", flag.ExitOnError)
	var port = flagSet.Uint("p", 0, "TCP port number")
	parseError := flagSet.Parse([]string {"-p", "50"})
	if parseError != nil {
		fmt.Printf("Failed to parse arguments: %s\n", parseError.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", flagSet.Name())
	fmt.Printf("- Parsed?: %s\n", flagSet.Parsed())
	fmt.Printf("- port: %d\n", *port)
	fmt.Printf("- Other stuff: [%d] %s\n", len(flagSet.Args()), flagSet.Args())
}
