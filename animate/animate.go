package animate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gosuri/uilive"
	"github.com/socialviolation/asciiban/ascii"
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

var blink = Sequence{
	Frames: []Frame{
		{
			Duration: 500 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteBlack),
			},
		}, {
			Duration: 500 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteWhite),
			},
		},
	},
}

var anaglygh = Sequence{
	XPadding: 2,
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

var glitch = Sequence{
	XPadding: 2,
	YPadding: 2,
	Frames: []Frame{
		{
			Duration: 3 * time.Second,
			Opts:     []ascii.BannerOption{},
		},
		{
			XOffset:  1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteLime),
				ascii.WithFont(ascii.FontTheEdge),
			},
		},
		{
			XOffset:  1,
			YOffset:  1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteMatrixGreen),
				ascii.WithFont(ascii.FontTheEdge),
			},
		},
		{
			XOffset:  -1,
			YOffset:  0,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteRed),
				ascii.WithFont(ascii.FontElite),
			},
		},
		{
			XOffset:  2,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteMintGreen),
				ascii.WithFont(ascii.FontElite),
			},
		},
		{
			XOffset:  0,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteIceBlue),
				ascii.WithFont(ascii.FontElite),
			},
		},
		{
			XOffset:  2,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteGoogle),
				ascii.WithFont(ascii.FontBloody),
			},
		},
		{
			XOffset:  0,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteRetroIcyPole),
				ascii.WithFont(ascii.FontBloody),
			},
		},
		{
			XOffset:  -2,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteRedBlack),
				ascii.WithFont(ascii.FontDOSRebel),
			},
		},
		{
			XOffset:  -3,
			YOffset:  -1,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PalettePatriot),
				ascii.WithFont(ascii.FontDOSRebel),
			},
		},
	},
}

var styleMap = map[string]Sequence{
	"anaglyph": anaglygh,
	"blink":    blink,
	"default":  blink,
	"glitch":   glitch,
}

func GetSequence(s string) Sequence {
	if val, ok := styleMap[s]; ok {
		return val
	}
	fmt.Println("Sequence not found, using default (glitch)")
	return GetSequence("default")
}

func Animate(ctx context.Context, seq Sequence, opts ...ascii.BannerOption) {
	frameIdx := 0
	writer := uilive.New()
	writer.Start()

	draw := func(f Frame) {
		writer.Flush()
		frameOpts := append(opts, f.Opts...)
		banner := ascii.Render(frameOpts...)
		banner = pad(seq.XPadding+f.XOffset, seq.YPadding+f.YOffset, banner)
		//fmt.Fprintln(writer, "frameIdx: "+strconv.Itoa(frameIdx))
		fmt.Fprintf(writer, banner)
	}

	draw(seq.Frames[frameIdx])
	for {
		select {
		case <-ctx.Done():
			writer.Stop()
			return
		case <-time.After(seq.Frames[frameIdx].Duration):
			frameIdx = (frameIdx + 1) % len(seq.Frames)
			draw(seq.Frames[frameIdx])
		}
	}
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
