//go:build !exclude

package ascii

import "fmt"

type Args struct {
	message    string
	font       string
	palette    Palette
	colourMode ColourMode
	trim       bool
	verbose    bool
}

type BannerOption func(*Args)

func WithMessage(m string) BannerOption {
	return func(args *Args) {
		args.message = m
	}
}

func WithFont(f string) BannerOption {
	return func(args *Args) {
		args.font = MatchFont(f)
	}
}

func WithPaletteName(p string) BannerOption {
	return func(args *Args) {
		args.palette = GetPalette(p)
	}
}
func WithPalette(p Palette) BannerOption {
	return func(args *Args) {
		args.palette = p
	}
}

func WithColourModeName(c string) BannerOption {
	return func(args *Args) {
		args.colourMode = GetColourMode(c)
	}
}

func WithColourMode(m ColourMode) BannerOption {
	return func(args *Args) {
		args.colourMode = m
	}
}

func WithTrim(trim bool) BannerOption {
	return func(args *Args) {
		args.trim = trim
	}
}

func WithVerbose(verbose bool) BannerOption {
	return func(args *Args) {
		args.verbose = verbose
	}
}

var (
	message = "asciiban"
)

func buildArgs(opts ...BannerOption) *Args {
	args := &Args{}
	for _, opt := range opts {
		opt(args)
	}
	if args.font == "" {
		args.font = "default"
	}
	if args.palette.IsEmpty() {
		args.palette = PaletteWhite
	}
	if args.message == "" {
		args.message = message
	}

	return args
}

func Print(opts ...BannerOption) {
	args := buildArgs(opts...)

	flf, err := loadFont(args.font)
	if err != nil {
		panic(err)
	}
	flf.Draw(*args)
}

func Random(opts ...BannerOption) {
	args := buildArgs(opts...)

	args.font = pickKeyFromMap(fontMap)
	args.palette = pickValueFromMap(PaletteMap)
	flf, err := loadFont(args.font)
	if err != nil {
		panic(err)
	}

	if args.verbose {
		fmt.Printf("font: %s, \nPalette: %s (%s)\n", args.font, args.palette.Name, args.palette.Key)
	}

	flf.Render(*args)
}
