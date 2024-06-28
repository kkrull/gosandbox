package rootCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: `cobrahello demonstrates how to use Cobra to parse command line arguments`,
	Run: func(cmd *cobra.Command, positionalArgs []string) {
		fmt.Println("rootCmd positional args:")
		for i, arg := range positionalArgs {
			fmt.Printf("- %d: %s\n", i, arg)
		}
	},
	Short:   "Demonstrates how to use Cobra",
	Use:     "cobrahello",
	Version: "0.0.1",
}

func init() {
	rootCmd.Flags().String("home-dir", "<default-value>", "some description of what this means")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
