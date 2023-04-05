package cli

import (
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/cprofiles"
	"os"

	"github.com/spf13/cobra"
)

var fillBg bool
var palette string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciiban",
	Short: "Generate ascii banners",
	Run: func(cmd *cobra.Command, args []string) {
		a := asciiban.DefaultArgs
		if len(args) != 0 {
			a.Message = args[0]
		}
		a.FillBg = fillBg
		a.Profile = cprofiles.Get(palette)
		asciiban.Print(a)
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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&fillBg, "fill-bg", "f", false, "Fill whitespace characters (doesn't look great in all fonts)")
	rootCmd.PersistentFlags().StringVarP(&palette, "palette", "p", "default", "Colour palette to use")
}
