package cmd

import (
	"os"

	"github.com/socialviolation/asciiban/ascii"

	"github.com/spf13/cobra"
)

var palette string
var font string
var mode string
var trim bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciiban",
	Short: "Generate ascii banner",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
		ascii.Print(a)
	},
}

// rootCmd represents the base command when called without any subcommands
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate ascii banner using random font & colours",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
		ascii.Random(a)
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
	rootCmd.PersistentFlags().BoolVarP(&trim, "trim", "t", true, "trim empty lines")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
}

func getArgs(args []string) ascii.Args {
	a := ascii.DefaultArgs
	if len(args) > 0 {
		a.message = args[0]
	}
	a.verbose = verbose
	var err error
	a.font = ascii.MatchFont(font)
	if err != nil {
		panic(err)
	}
	a.Palette = ascii.GetPalette(palette)
	if mode != "" {
		a.colourMode = ascii.GetColourMode(mode)
	}

	return a
}
