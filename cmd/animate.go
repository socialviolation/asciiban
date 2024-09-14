package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/socialviolation/asciiban/animate"
	"github.com/spf13/cobra"
)

var sequence string

var animateCmd = &cobra.Command{
	Use:   "animate",
	Short: "Animate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getOpts(args)
		ctx := context.Background()
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		seq := animate.GetSequence(sequence)
		animate.Animate(ctx, seq, a...)
	},
}

func init() {
	rootCmd.AddCommand(animateCmd)
	animateCmd.PersistentFlags().StringVarP(&sequence, "sequence", "s", "default", "Animation sequence")
}
