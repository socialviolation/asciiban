package cli

import (
	"github.com/socialviolation/asciiban"
	"github.com/socialviolation/asciiban/fonts"
	"github.com/spf13/cobra"
)

var message string

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		asciiban.Print(asciiban.Args{
			Message: message,
			Profile: asciiban.MintGreen,
			Font:    fonts.ANSIShadow,
		})
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Banner to print")
	_ = printCmd.MarkFlagRequired("message")
}
