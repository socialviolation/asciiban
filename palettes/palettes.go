package palettes

type ColourMode int64

const (
	Nil ColourMode = iota
	Single
	Alternate
	VerticalGradient
	HorizontalGradient
)

type Palette struct {
	// List of Hex Codes to use in scheme
	Colours    []string
	ColourMode ColourMode
}

func (c *Palette) IsEmpty() bool {
	return c.Colours == nil
}

// https://colorhunt.co/palettes/neon
var (
	MintGreen = Palette{
		Colours:    []string{"CCFFBD", "7ECA9C", "4b795d"},
		ColourMode: VerticalGradient,
	}
	RetroIcyPole = Palette{
		Colours:    []string{"F38181", "FCE38A", "EAFFD0", "95E1D3"},
		ColourMode: VerticalGradient,
	}
	Desert = Palette{
		Colours:    []string{"FFACAC", "FFBFA9", "FFEBB4", "FBFFB1"},
		ColourMode: VerticalGradient,
	}
	IceBlue = Palette{
		Colours:    []string{"E3FDFD", "CBF1F5", "A6E3E9", "71C9CE"},
		ColourMode: VerticalGradient,
	}
	SwampGreen = Palette{
		Colours:    []string{"DDFFBB", "C7E9B0", "B3C99C"},
		ColourMode: VerticalGradient,
	}
	BogGreen = Palette{
		Colours:    []string{"7DCE13", "70b911", "64a40f"},
		ColourMode: VerticalGradient,
	}
	MatrixGreen = Palette{
		Colours:    []string{"00FF41", "00cc34", "009927"},
		ColourMode: VerticalGradient,
	}
	Google = Palette{
		Colours:    []string{"4285F4", "DB4437", "F4B400", "0F9D58"},
		ColourMode: VerticalGradient,
	}
	White = Palette{
		Colours:    []string{"FFFFFF"},
		ColourMode: Single,
	}
	Red = Palette{
		Colours:    []string{"ff0000"},
		ColourMode: Single,
	}
	Green = Palette{
		Colours:    []string{"008000"},
		ColourMode: Single,
	}
	Blue = Palette{
		Colours:    []string{"0000ff"},
		ColourMode: Single,
	}
	Yellow = Palette{
		Colours:    []string{"ffff00"},
		ColourMode: Single,
	}
	Purple = Palette{
		Colours:    []string{"A020F0"},
		ColourMode: Single,
	}
	RedBlack = Palette{
		Colours:    []string{"ff0000", "36454F"},
		ColourMode: Alternate,
	}
	RedOrange = Palette{
		Colours:    []string{"ff0000", "FF5733"},
		ColourMode: Alternate,
	}
	Pizza = Palette{
		Colours:    []string{"008c45", "f4f5f0", "cd212a"},
		ColourMode: HorizontalGradient,
	}
)

var ProfileMap = map[string]Palette{
	"mint":       MintGreen,
	"matrix":     MatrixGreen,
	"bog":        BogGreen,
	"swamp":      SwampGreen,
	"ice":        IceBlue,
	"desert":     Desert,
	"retro":      RetroIcyPole,
	"google":     Google,
	"default":    White,
	"red":        Red,
	"green":      Green,
	"blue":       Blue,
	"yellow":     Yellow,
	"purple":     Purple,
	"red-black":  RedBlack,
	"red-orange": RedOrange,
	"pizza":      Pizza,
}

func Get(p string) Palette {
	if val, ok := ProfileMap[p]; ok {
		return val
	}
	return White
}

func GetMode(p string) ColourMode {
	switch p {
	case "s":
		fallthrough
	case "single":
		return Single
	case "a":
		fallthrough
	case "alt":
		fallthrough
	case "alternating":
		return Alternate
	case "v":
		fallthrough
	case "vert":
		fallthrough
	case "vertical":
		return VerticalGradient
	case "h":
		fallthrough
	case "horiz":
		fallthrough
	case "horizontal":
		return HorizontalGradient
	}
	return Nil
}
