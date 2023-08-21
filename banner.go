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
	if args.Font == nil && args.Palette.IsEmpty() {
		args.Font, _ = GetFont("default")
		args.Palette = White
	}
	if args.Font == nil {
		args.Font, _ = GetFont("default")
	}
	if args.Palette.IsEmpty() {
		args.Palette = White
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
