package cli

import (
	"github.com/socialviolation/asciiban"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciiban",
	Short: "Generate ascii banners",
	Run: func(cmd *cobra.Command, args []string) {
		asciiban.Print(asciiban.DefaultArgs)
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
