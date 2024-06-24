package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Arguments: %v\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("- %d: %s\n", i, arg)
	}

	flagSet := flag.NewFlagSet("custom program", flag.ContinueOnError)
	port := flagSet.Uint("p", 0, "TCP port number")
	parseError := flagSet.Parse(os.Args[1:])
	if parseError == flag.ErrHelp {
		fmt.Printf("Asked for help")
		os.Exit(0)
	} else if parseError != nil {
		fmt.Printf("Failed to parse arguments: %s\n", parseError.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", flagSet.Name())
	fmt.Printf("- Parsed?: %s\n", flagSet.Parsed())
	fmt.Printf("- port: %d\n", *port)
	fmt.Printf("- Other stuff: [%d] %s\n", len(flagSet.Args()), flagSet.Args())
}
