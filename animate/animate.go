package animate

import (
	"context"
	"fmt"
	"io"
	"math"
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
	YPadding: 3,
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
			YOffset:  0,
			Duration: 100 * time.Millisecond,
			Opts: []ascii.BannerOption{
				ascii.WithPalette(ascii.PaletteRed),
				ascii.WithFont(ascii.FontBloody),
			},
		},
		{
			XOffset:  0,
			YOffset:  1,
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
	"3d":       anaglygh,
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
	maxLines := 0
	for _, f := range seq.Frames {
		frameOpts := append(opts, f.Opts...)
		banner := ascii.Render(frameOpts...)
		lh := strings.Count(banner, "\n") + (seq.YPadding * 2) + f.YOffset
		if lh > maxLines {
			maxLines = lh
		}
	}

	draw := func(w io.Writer, f Frame) {
		frameOpts := append(opts, f.Opts...)
		banner := ascii.Render(frameOpts...)
		banner = pad(seq.XPadding, seq.YPadding, f.XOffset, f.YOffset, banner)
		lh := strings.Count(banner, "\n")
		buffer := int(math.Max(float64(maxLines-lh), 0))
		if lh > 0 {
			banner += strings.Repeat("\n", buffer)
		}
		_, err := fmt.Fprintf(w, banner)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	writer := uilive.New()
	writer.Start()
	defer func() {
		err := writer.Flush()
		if err != nil {
			fmt.Println(err)
		}
		writer.Stop()
	}()

	frameIdx := 0
	ticker := time.NewTicker(seq.Frames[frameIdx].Duration)

	quit := make(chan bool)
	draw(writer, seq.Frames[frameIdx])
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				quit <- true
				fmt.Println("context cancelled")
				return
			case <-ticker.C:
				frameIdx = (frameIdx + 1) % len(seq.Frames)
				draw(writer, seq.Frames[frameIdx])
				writer.Flush()
				ticker.Stop()
				ticker = time.NewTicker(seq.Frames[frameIdx].Duration)
			}
		}
	}()
	<-quit
}

func pad(xPad int, yPad int, xOff int, yOff int, rdr string) string {
	if xPad+xOff > 0 {
		xStr := strings.Repeat(" ", xPad+xOff)
		rdr = xStr + rdr
		rdr = strings.ReplaceAll(rdr, "\n", "\n"+xStr)
	}

	if yPad+yOff >= 0 && yPad-yOff >= 0 {
		yPre := strings.Repeat("\n", yPad+yOff)
		ySuf := strings.Repeat("\n", yPad-yOff)
		rdr = yPre + rdr + ySuf
	}

	return rdr
}
