package cmd

import (
	"fmt"

	"github.com/socialviolation/asciiban/ascii"
	"github.com/spf13/cobra"
)

// palettesCmd represents the palettes command
var palettesCmd = &cobra.Command{
	Use:   "palettes",
	Short: "Subcommands show info for available palettes",
}

// palettesListCmd represents the list command
var palettesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available colour palettes",
	Run: func(cmd *cobra.Command, args []string) {
		for p, _ := range ascii.PaletteMap {
			fmt.Println(p)
		}
	},
}

// palettesListCmd represents the list command
var palettesTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test all available colour palettes",
	Run: func(cmd *cobra.Command, args []string) {
		a := getArgs(args)
		for p, _ := range ascii.PaletteMap {
			fmt.Println(p)
			a.Palette = ascii.GetPalette(p)
			ascii.Print(a)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(palettesCmd)
	palettesCmd.AddCommand(palettesListCmd)
	palettesCmd.AddCommand(palettesTestCmd)
}
