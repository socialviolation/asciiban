package cmd

import (
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/palettes"
	"os"

	"github.com/spf13/cobra"
)

var palette string
var font string
var mode string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciiban",
	Short: "Generate ascii banners",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
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
	rootCmd.PersistentFlags().StringVarP(&palette, "palette", "p", "default", "Colour palette to use")
	rootCmd.PersistentFlags().StringVarP(&font, "font", "f", "ansishadow", "Colour palette to use")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "Palette Colour Mode (simple | alternating | vertical | horizontal)")
}

func getArgs(args []string) asciiban.Args {
	a := asciiban.DefaultArgs
	if len(args) > 0 {
		a.Message = args[0]
	}
	a.Font = fontspack.Get(font)
	a.Palette = palettes.Get(palette)
	a.ColourMode = palettes.GetMode(mode)

	return a
}
