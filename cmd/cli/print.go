package cli

import (
	"github.com/socialviolation/asciiban"
	"github.com/spf13/cobra"
)

var message string
var fillBg bool
var palette string

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print a banner",
	Run: func(cmd *cobra.Command, args []string) {
		a := asciiban.DefaultArgs
		a.Message = message
		a.FillBg = fillBg
		a.Profile = asciiban.GetProfile(palette)
		asciiban.Print(a)
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message to print")
	_ = printCmd.MarkFlagRequired("message")
	printCmd.PersistentFlags().BoolVarP(&fillBg, "fill-bg", "f", false, "Fill whitespace characters (doesn't look great in all fonts)")
	printCmd.PersistentFlags().StringVarP(&palette, "palette", "p", "default", "Colour palette to use")
}
