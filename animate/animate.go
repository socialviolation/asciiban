package animate

import (
	"context"
	"fmt"
	"github.com/gosuri/uilive"
	"github.com/socialviolation/asciiban/ascii"
	"strings"
	"time"
)

type Sequence struct {
	XPadding int
	YPadding int
	Frames   []Frame
}

type Frame struct {
	XOffset  int
	YOffset  int
	Duration time.Duration
	Opts     []ascii.BannerOption
}

var anaglygh = Sequence{
	XPadding: 3,
	YPadding: 0,
	Frames: []Frame{
		{
			XOffset:  -1,
			Duration: 30 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteRed),
			},
		},
		{
			XOffset:  1,
			Duration: 30 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteCyan),
			},
		},
	},
}

func Anaglyph(ctx context.Context, opts ...ascii.BannerOption) {
	frame := 0

	writer := uilive.New()
	writer.Start()

	for {
		select {
		case <-ctx.Done():
			writer.Stop()
			return
		default:
			writer.Flush()
			frame = (frame + 1) % len(anaglygh.Frames)
			frameOpts := append(opts, anaglygh.Frames[frame].Opts...)
			rsr := ascii.Render(frameOpts...)
			rsr = pad(anaglygh.XPadding+anaglygh.Frames[frame].XOffset, anaglygh.YPadding, rsr)
			fmt.Fprintf(writer, "\n"+rsr)
			time.Sleep(anaglygh.Frames[frame].Duration)
		}
	}
}

func Glitch() {

}

func pad(x int, y int, rdr string) string {
	if x > 0 {
		xpad := strings.Repeat(" ", x)
		rdr = xpad + rdr
		rdr = strings.ReplaceAll(rdr, "\n", "\n"+xpad)
	}

	if y > 0 {
		ypad := strings.Repeat("\n", y)
		rdr = ypad + rdr + ypad
	}

	return rdr
}