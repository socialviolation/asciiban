package cmd

import (
	"fmt"
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/fonts"
	"github.com/socialviolation/asciiban/palettes"

	"github.com/spf13/cobra"
)

// fontsCmd represents the fonts command
var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "A brief description of your command",
}

// listCmd represents the list command
var fontsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available fonts",
	Run: func(cmd *cobra.Command, args []string) {
		for f, _ := range fonts.FontMap {
			fmt.Println(f)
		}
	},
}

// listCmd represents the list command
var fontsTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all available fonts",
	Run: func(cmd *cobra.Command, args []string) {
		for f, _ := range fonts.FontMap {
			fmt.Println(f)
			a := asciiban.DefaultArgs
			if len(args) > 0 {
				a.Message = args[0]
			}
			a.Font = fonts.Get(f)
			a.FillBg = fillBg
			a.Profile = palettes.Get(palette)
			asciiban.Print(a)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(fontsListCmd)
	fontsCmd.AddCommand(fontsTestCmd)
}
