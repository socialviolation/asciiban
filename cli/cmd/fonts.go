package cmd

import (
	"fmt"
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/fontpack"
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
	Short: "List all available fontpack",
	Run: func(cmd *cobra.Command, args []string) {
		for f, _ := range fontpack.FontMap {
			fmt.Println(f)
		}
	},
}

// listCmd represents the list command
var fontsTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all available fontpack",
	Run: func(cmd *cobra.Command, args []string) {
		keys := make([]string, 0, len(fontpack.FontMap))

		for k := range fontpack.FontMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		a := getArgs(args)
		for _, k := range keys {
			fmt.Println(k)
			a.Font = fontpack.Get(k)
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
