package cmd

import (
	"context"
	"github.com/socialviolation/asciiban/animate"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var animateCmd = &cobra.Command{
	Use:   "animate",
	Short: "Animate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getOpts(args)
		ctx := context.Background()
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		animate.Glitch(ctx, a...)
	},
}

func init() {
	rootCmd.AddCommand(animateCmd)
}
