package cmd

import (
	"fmt"
	"github.com/socialviolation/asciiban"
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
		for f, _ := range asciiban.FontMap {
			fmt.Println(f)
		}
	},
}

// listCmd represents the list command
var fontsTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all available fonts",
	Run: func(cmd *cobra.Command, args []string) {
		keys := make([]string, 0, len(asciiban.FontMap))

		for k := range asciiban.FontMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		a := getArgs(args)
		for _, k := range keys {
			fmt.Println(k)
			a.Font = asciiban.GetFont(k)
			asciiban.Print(a)
			fmt.Println()
		}
	},
}

// listCmd represents the list command
var parseTestCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse font",
	Run: func(cmd *cobra.Command, args []string) {
		keys := make([]string, 0, len(asciiban.FontMap))
		for k := range asciiban.FontMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		a := getArgs(args)
		for _, k := range keys {
			fmt.Println(k)
			a.Font = asciiban.GetFont(k)
			f, err := asciiban.ParseFlf(k, a.Font)
			if err != nil {
				fmt.Println("error parsing font " + k)
				continue
			}
			fmt.Print(f.Render(a.Message))

			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(fontsListCmd)
	fontsCmd.AddCommand(fontsTestCmd)
	fontsCmd.AddCommand(parseTestCmd)
}
