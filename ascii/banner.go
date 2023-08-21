//go:build !exclude

package ascii

import "fmt"

type Args struct {
	Message    string
	FontName   string
	Palette    Palette
	ColourMode ColourMode
	Trim       bool
	Verbose    bool
}

var DefaultArgs = Args{
	Message:  "ascii banner",
	FontName: "",
	Palette:  PaletteDefault,
}

func Print(args Args) {
	if args.FontName == "" {
		args.FontName = "default"
	}
	if args.Palette.IsEmpty() {
		args.Palette = White
	}

	flf, err := loadFont(args.FontName)
	if err != nil {
		panic(err)
	}
	flf.Render(args)
}

func Random(args Args) {
	args.FontName = pickKeyFromMap(fontMap)
	args.Palette = pickValueFromMap(ProfileMap)
	flf, err := loadFont(args.FontName)
	if err != nil {
		panic(err)
	}

	if args.Verbose {
		fmt.Printf("Font: %s, \nPalette: %s (%s)\n", args.FontName, args.Palette.Name, args.Palette.Key)
	}
	flf.Render(args)
}
