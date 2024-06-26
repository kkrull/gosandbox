package rootCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-hello",
	Short: "Demonstrates how to use Cobra",
	Long: `Demonstrates how to use Cobra to parse Command Line Arguments:
Usage: cobrahello [-h|--help] [-t|--toggle]`,
	Run: func(cmd *cobra.Command, positionalArgs []string) {
		fmt.Println("rootCmd positional args:")
		for i, arg := range positionalArgs {
			fmt.Printf("- %d: %s\n", i, arg)
		}
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Define global (persistent) and local flags and configuration settings.
func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-hello.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
