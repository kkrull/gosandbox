package rootCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: `Demonstrates how to use Cobra to parse Command Line Arguments:
Usage: cobrahello [-h|--help]`,
	Run: func(cmd *cobra.Command, positionalArgs []string) {
		fmt.Println("rootCmd positional args:")
		for i, arg := range positionalArgs {
			fmt.Printf("- %d: %s\n", i, arg)
		}
	},
	Short: "Demonstrates how to use Cobra",
	Use:   "cobra-hello",
	Version: "0.0.1",
}

func init() {
	// TODO KDK: Add a local flag and see that it gets parsed
	// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#local-flags
	rootCmd.Flags().String("home-dir", "<default-value>", "some description of what this means")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
