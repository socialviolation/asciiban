//go:build !exclude

//go:generate go run gen.go
package asciiban

type Args struct {
	Message    string
	Font       *Font
	Palette    Palette
	ColourMode ColourMode
	Trim       bool
}

var DefaultArgs Args = Args{
	Message: "ascii banner",
	Font:    nil,
	Palette: PaletteDefault,
}

func Print(args Args) {
	if args.Font == nil {
		if args.ColourMode == modeStarsNStripes || args.Palette.ColourMode == modeStarsNStripes {
			args.Palette = PaletteMurica
			args.ColourMode = modeStarsNStripes
			args.Font, _ = GetFont("usaflag")
		} else {
			args.Font, _ = GetFont("default")
		}
	}
	if args.Palette.IsEmpty() {
		args.Palette = White
	}
	if args.ColourMode != modeNil {
		args.Palette.ColourMode = args.ColourMode
	}

	args.Font.Render(args)
}

func Random(args Args) {
	var err error
	args.Font, err = GetFont(pick(FontMap))
	if err != nil {
		panic(err)
	}
	args.Palette = pick(ProfileMap)
	args.ColourMode = modeNil

	args.Font.Render(args)
}
