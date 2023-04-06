package palettes

type ColourMode int64

const (
	VerticalGradient ColourMode = iota
	HorizontalGradient
	Single
	Alternate
)

type Palette struct {
	// List of Hex Codes to use in scheme
	Colours    []string
	RenderType ColourMode
}

func (c *Palette) IsEmpty() bool {
	return c.Colours == nil
}

// https://colorhunt.co/palettes/neon
var (
	MintGreen = Palette{
		Colours:    []string{"CCFFBD", "7ECA9C", "4b795d"},
		RenderType: VerticalGradient,
	}
	RetroIcyPole = Palette{
		Colours:    []string{"F38181", "FCE38A", "EAFFD0", "95E1D3"},
		RenderType: VerticalGradient,
	}
	Desert = Palette{
		Colours:    []string{"FFACAC", "FFBFA9", "FFEBB4", "FBFFB1"},
		RenderType: VerticalGradient,
	}
	IceBlue = Palette{
		Colours:    []string{"E3FDFD", "CBF1F5", "A6E3E9", "71C9CE"},
		RenderType: VerticalGradient,
	}
	SwampGreen = Palette{
		Colours:    []string{"DDFFBB", "C7E9B0", "B3C99C"},
		RenderType: VerticalGradient,
	}
	BogGreen = Palette{
		Colours:    []string{"7DCE13", "70b911", "64a40f"},
		RenderType: VerticalGradient,
	}
	MatrixGreen = Palette{
		Colours:    []string{"00FF41", "00cc34", "009927"},
		RenderType: VerticalGradient,
	}
	Google = Palette{
		Colours:    []string{"4285F4", "DB4437", "F4B400", "0F9D58"},
		RenderType: VerticalGradient,
	}
	White = Palette{
		Colours:    []string{"FFFFFF"},
		RenderType: Single,
	}
	Red = Palette{
		Colours:    []string{"ff0000"},
		RenderType: Single,
	}
	Green = Palette{
		Colours:    []string{"008000"},
		RenderType: Single,
	}
	Blue = Palette{
		Colours:    []string{"0000ff"},
		RenderType: Single,
	}
	Yellow = Palette{
		Colours:    []string{"ffff00"},
		RenderType: Single,
	}
	Purple = Palette{
		Colours:    []string{"A020F0"},
		RenderType: Single,
	}
)

var ProfileMap = map[string]Palette{
	"mint":    MintGreen,
	"matrix":  MatrixGreen,
	"bog":     BogGreen,
	"swamp":   SwampGreen,
	"ice":     IceBlue,
	"desert":  Desert,
	"retro":   RetroIcyPole,
	"google":  Google,
	"default": White,
	"red":     Red,
	"green":   Green,
	"blue":    Blue,
	"yellow":  Yellow,
	"purple":  Purple,
}

func Get(p string) Palette {
	if val, ok := ProfileMap[p]; ok {
		return val
	}
	return White
}
