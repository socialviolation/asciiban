package cmd

import (
	"github.com/socialviolation/asciiban"
	"os"

	"github.com/spf13/cobra"
)

var palette string
var font string
var mode string
var trim bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciicli",
	Short: "Generate ascii banner",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
		asciiban.Print(a)
	},
}

// rootCmd represents the base command when called without any subcommands
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate ascii banner using random font & colours",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
		asciiban.Random(a)
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
	rootCmd.AddCommand(randomCmd)
	rootCmd.PersistentFlags().StringVarP(&palette, "palette", "p", "default", "Colour palette to use")
	rootCmd.PersistentFlags().StringVarP(&font, "font", "f", "ansishadow", "Colour palette to use")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "Palette Colour Mode (simple | alternating | vertical | horizontal)")
	rootCmd.PersistentFlags().BoolVarP(&trim, "trim", "t", true, "Trim empty lines")
}

func getArgs(args []string) asciiban.Args {
	a := asciiban.DefaultArgs
	if len(args) > 0 {
		a.Message = args[0]
	}
	var err error
	a.Font, err = asciiban.GetFont(font)
	if err != nil {
		panic(err)
	}
	a.Palette = asciiban.GetPalette(palette)
	if mode != "" {
		a.ColourMode = asciiban.GetColourMode(mode)
	} else {
		a.ColourMode = a.Palette.ColourMode
	}

	return a
}
