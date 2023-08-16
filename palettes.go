//go:build !exclude

package asciiban

import "fmt"

type ColourMode int64

const (
	modeNil ColourMode = iota
	modeSingle
	modeAlternate
	modeVerticalGradient
	modeHorizontalGradient
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
	PaletteMintGreen = Palette{
		Colours:    []string{"CCFFBD", "7ECA9C", "4b795d"},
		ColourMode: modeVerticalGradient,
	}
	PaletteRetroIcyPole = Palette{
		Colours:    []string{"F38181", "FCE38A", "EAFFD0", "95E1D3"},
		ColourMode: modeVerticalGradient,
	}
	PaletteDesert = Palette{
		Colours:    []string{"FFACAC", "FFBFA9", "FFEBB4", "FBFFB1"},
		ColourMode: modeVerticalGradient,
	}
	PaletteIceBlue = Palette{
		Colours:    []string{"E3FDFD", "CBF1F5", "A6E3E9", "71C9CE"},
		ColourMode: modeVerticalGradient,
	}
	PaletteSwampGreen = Palette{
		Colours:    []string{"DDFFBB", "C7E9B0", "B3C99C"},
		ColourMode: modeVerticalGradient,
	}
	PaletteBogGreen = Palette{
		Colours:    []string{"7DCE13", "70b911", "64a40f"},
		ColourMode: modeVerticalGradient,
	}
	PaletteMatrixGreen = Palette{
		Colours:    []string{"00FF41", "00cc34", "009927"},
		ColourMode: modeVerticalGradient,
	}
	PaletteGoogle = Palette{
		Colours:    []string{"4285F4", "DB4437", "F4B400", "0F9D58"},
		ColourMode: modeVerticalGradient,
	}
	White = Palette{
		Colours:    []string{"FFFFFF"},
		ColourMode: modeSingle,
	}
	PaletteRed = Palette{
		Colours:    []string{"ff0000"},
		ColourMode: modeSingle,
	}
	PaletteGreen = Palette{
		Colours:    []string{"008000"},
		ColourMode: modeSingle,
	}
	PaletteBlue = Palette{
		Colours:    []string{"0000ff"},
		ColourMode: modeSingle,
	}
	PaletteYellow = Palette{
		Colours:    []string{"ffff00"},
		ColourMode: modeSingle,
	}
	PalettePurple = Palette{
		Colours:    []string{"A020F0"},
		ColourMode: modeSingle,
	}
	PaletteRedBlack = Palette{
		Colours:    []string{"ff0000", "36454F"},
		ColourMode: modeAlternate,
	}
	PaletteRedOrange = Palette{
		Colours:    []string{"ff0000", "FF5733"},
		ColourMode: modeAlternate,
	}
	PalettePizza = Palette{
		Colours:    []string{"008c45", "f4f5f0", "cd212a"},
		ColourMode: modeHorizontalGradient,
	}
	PaletteMurica = Palette{
		Colours:    []string{"BB133E", "FFFFFF", "004594"},
		ColourMode: modeAlternate,
	}
)

var PaletteDefault = White

var ProfileMap = map[string]Palette{
	"mint":       PaletteMintGreen,
	"matrix":     PaletteMatrixGreen,
	"bog":        PaletteBogGreen,
	"swamp":      PaletteSwampGreen,
	"ice":        PaletteIceBlue,
	"desert":     PaletteDesert,
	"retro":      PaletteRetroIcyPole,
	"google":     PaletteGoogle,
	"default":    PaletteDefault,
	"red":        PaletteRed,
	"green":      PaletteGreen,
	"blue":       PaletteBlue,
	"yellow":     PaletteYellow,
	"purple":     PalettePurple,
	"red-black":  PaletteRedBlack,
	"red-orange": PaletteRedOrange,
	"pizza":      PalettePizza,
	"murica":     PaletteMurica,
}

func GetPalette(p string) Palette {
	if val, ok := ProfileMap[p]; ok {
		return val
	}
	fmt.Println("Palette not found, using default palette")
	return GetPalette("default")
}

func GetPaletteMode(p string) ColourMode {
	switch p {
	case "s":
		fallthrough
	case "single":
		return modeSingle
	case "a":
		fallthrough
	case "alt":
		fallthrough
	case "alternating":
		return modeAlternate
	case "v":
		fallthrough
	case "vert":
		fallthrough
	case "vertical":
		return modeVerticalGradient
	case "h":
		fallthrough
	case "horiz":
		fallthrough
	case "horizontal":
		return modeHorizontalGradient
	}
	return modeNil
}
