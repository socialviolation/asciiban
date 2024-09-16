package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/socialviolation/asciiban/animate"
	"github.com/spf13/cobra"
	//_ "net/http/pprof"
)

var sequence string
var duration int

var animateCmd = &cobra.Command{
	Use:   "animate",
	Short: "Animate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := getOpts(args)
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		defer cancel()
		if duration > 0 {
			ctx, cancel = context.WithTimeout(ctx, time.Duration(duration)*time.Second)
			defer cancel()
		}
		animate.Animate(ctx, animate.GetSequence(sequence), a...)
		cancel()
	},
}

func init() {
	rootCmd.AddCommand(animateCmd)
	animateCmd.PersistentFlags().StringVarP(&sequence, "sequence", "s", "default", "Animation sequence")
	animateCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 0, "Animation duration")
}
