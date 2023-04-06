package cmd

import (
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/fonts"
	"github.com/socialviolation/asciiban/palettes"
	"os"

	"github.com/spf13/cobra"
)

var fillBg bool
var palette string
var font string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciiban",
	Short: "Generate ascii banners",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := asciiban.DefaultArgs
		a.Message = args[0]
		a.Font = fonts.Get(font)
		a.FillBg = fillBg
		a.Profile = palettes.Get(palette)
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
	rootCmd.PersistentFlags().BoolVarP(&fillBg, "background", "b", false, "Fill whitespace characters (doesn't look great in all fonts)")
	rootCmd.PersistentFlags().StringVarP(&palette, "palette", "p", "default", "Colour palette to use")
	rootCmd.PersistentFlags().StringVarP(&font, "font", "f", "ansishadow", "Colour palette to use")
}