package rootCmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	cmdOutWriter io.Writer
	homeDirFlag  *string
	rootCmd      = &cobra.Command{
		Long: `cobrahello demonstrates how to use Cobra to parse command line arguments`,
		Run: func(cmd *cobra.Command, positionalArgs []string) {
			fmt.Fprintf(cmdOutWriter, "homeDir: %s\n", *homeDirFlag)

			fmt.Fprintln(cmdOutWriter, "rootCmd positional args:")
			for i, arg := range positionalArgs {
				fmt.Fprintf(cmdOutWriter, "- %d: %s\n", i, arg)
			}
		},
		Short:   "Demonstrates how to use Cobra",
		Use:     "cobrahello",
		Version: "0.0.1",
	}
)

func init() {
	homeDirFlag = rootCmd.Flags().String(
		"home-dir",
		"<default-value>",
		"some description of what this means",
	)
}

func Execute(args []string, errWriter io.Writer, outWriter io.Writer) error {
	cmdOutWriter = outWriter

	rootCmd.SetArgs(args)
	rootCmd.SetErr(errWriter)
	rootCmd.SetOut(outWriter)
	return rootCmd.Execute()
}
