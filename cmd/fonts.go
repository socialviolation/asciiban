package cmd

import (
	"fmt"
	"github.com/socialviolation/asciiban/ascii"
	"sort"

	"github.com/spf13/cobra"
)

// fontsCmd represents the fontpack command
var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "Subcommands show info for available fonts",
}

// listCmd represents the list command
var fontsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available fonts",
	Run: func(cmd *cobra.Command, args []string) {
		fonts := ascii.GetFonts()
		for _, f := range fonts {
			fmt.Println(f)
		}
	},
}

// listCmd represents the list command
var fontsTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all available fonts",
	Run: func(cmd *cobra.Command, args []string) {
		fonts := ascii.GetFonts()
		sort.Strings(fonts)

		a := getArgs(args)
		for _, k := range fonts {
			fmt.Println(k)
			a.Font = k
			ascii.Print(a)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(fontsListCmd)
	fontsCmd.AddCommand(fontsTestCmd)
}
