//go:build !exclude

package ascii

import "fmt"

type Args struct {
	Message    string
	Font       string
	Palette    Palette
	ColourMode ColourMode
	Trim       bool
	Verbose    bool
}

var DefaultArgs = Args{
	Message: "ascii banner",
	Font:    FontANSIShadow,
	Palette: PaletteDefault,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = "default"
	}
	if args.Palette.IsEmpty() {
		args.Palette = White
	}

	flf, err := loadFont(args.Font)
	if err != nil {
		panic(err)
	}
	flf.Render(args)
}

func Random(args Args) {
	args.Font = pickKeyFromMap(fontMap)
	args.Palette = pickValueFromMap(ProfileMap)
	flf, err := loadFont(args.Font)
	if err != nil {
		panic(err)
	}

	if args.Verbose {
		fmt.Printf("Font: %s, \nPalette: %s (%s)\n", args.Font, args.Palette.Name, args.Palette.Key)
	}
	flf.Render(args)
}
