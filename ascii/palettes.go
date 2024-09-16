//go:build !exclude

package ascii

import "fmt"

type ColourMode int64

const (
	modeNil ColourMode = iota
	modeSingle
	modeAlternate
	modeVerticalGradient
	modeHorizontalGradient
	modeLetter
	modePatriot
)

type Palette struct {
	Name       string
	Key        string
	Colours    []string // List of Hex Codes to use in scheme
	ColourMode ColourMode
}

func (c *Palette) IsEmpty() bool {
	return c.ColourMode == modeNil
}

// https://colorhunt.co/palettes/neon
var (
	PaletteMintGreen = Palette{
		Name:       "Mint Green",
		Key:        "mint",
		Colours:    []string{"CCFFBD", "7ECA9C", "4b795d"},
		ColourMode: modeVerticalGradient,
	}
	PaletteRetroIcyPole = Palette{
		Name:       "Retro Icy Pole",
		Key:        "retro",
		Colours:    []string{"F38181", "FCE38A", "EAFFD0", "95E1D3"},
		ColourMode: modeVerticalGradient,
	}
	PaletteDesert = Palette{
		Name:       "Dessert",
		Key:        "desert",
		Colours:    []string{"FFACAC", "FFBFA9", "FFEBB4", "FBFFB1"},
		ColourMode: modeVerticalGradient,
	}
	PaletteIceBlue = Palette{
		Name:       "Ice Blue",
		Key:        "ice",
		Colours:    []string{"E3FDFD", "CBF1F5", "A6E3E9", "71C9CE"},
		ColourMode: modeVerticalGradient,
	}
	PaletteSwampGreen = Palette{
		Name:       "Swamp Green",
		Key:        "swamp",
		Colours:    []string{"DDFFBB", "C7E9B0", "B3C99C"},
		ColourMode: modeVerticalGradient,
	}
	PaletteBogGreen = Palette{
		Name:       "BOG™",
		Key:        "bog",
		Colours:    []string{"7DCE13", "70b911", "64a40f"},
		ColourMode: modeVerticalGradient,
	}
	PaletteMatrixGreen = Palette{
		Name:       "Matrix",
		Key:        "matrix",
		Colours:    []string{"00FF41", "00cc34", "009927"},
		ColourMode: modeVerticalGradient,
	}
	PaletteGoogle = Palette{
		Name:       "Google Theme",
		Key:        "google",
		Colours:    []string{"4285F4", "DB4437", "F4B400", "4285F4", "0F9D58", "DB4437"},
		ColourMode: modeLetter,
	}
	PaletteWhite = Palette{
		Name:       "PaletteWhite",
		Key:        "white",
		Colours:    []string{"FFFFFF"},
		ColourMode: modeSingle,
	}
	PaletteBlack = Palette{
		Name:       "PaletteBlack",
		Key:        "white",
		Colours:    []string{"000000"},
		ColourMode: modeSingle,
	}
	PaletteRed = Palette{
		Name:       "Red",
		Key:        "red",
		Colours:    []string{"ff0000"},
		ColourMode: modeSingle,
	}
	PaletteCyan = Palette{
		Name:       "Cyan",
		Key:        "cyan",
		Colours:    []string{"00FFFF"},
		ColourMode: modeSingle,
	}
	PaletteGreen = Palette{
		Name:       "Green",
		Key:        "green",
		Colours:    []string{"008000"},
		ColourMode: modeSingle,
	}
	PaletteLime = Palette{
		Name:       "Lime",
		Key:        "lime",
		Colours:    []string{"66ff00"},
		ColourMode: modeSingle,
	}
	PaletteBlue = Palette{
		Name:       "Blue",
		Key:        "blue",
		Colours:    []string{"0000ff"},
		ColourMode: modeSingle,
	}
	PaletteYellow = Palette{
		Name:       "Yellow",
		Key:        "yellow",
		Colours:    []string{"ffff00"},
		ColourMode: modeSingle,
	}
	PalettePurple = Palette{
		Name:       "Purple",
		Key:        "purple",
		Colours:    []string{"A020F0"},
		ColourMode: modeSingle,
	}
	PaletteRedBlack = Palette{
		Name:       "Red-Black",
		Key:        "red-black",
		Colours:    []string{"ff0000", "36454F"},
		ColourMode: modeAlternate,
	}
	PaletteRedOrange = Palette{
		Name:       "Red Orange",
		Key:        "red-orange",
		Colours:    []string{"ff0000", "FF5733"},
		ColourMode: modeAlternate,
	}
	PalettePizza = Palette{
		Name:       "Pizza Mode",
		Key:        "pizza",
		Colours:    []string{"008c45", "f4f5f0", "cd212a"},
		ColourMode: modeHorizontalGradient,
	}
	PalettePatriot = Palette{
		Name:       "Patriot Mode",
		Key:        "patriot",
		Colours:    []string{"ff0000", "FFFFFF", "003472"},
		ColourMode: modePatriot,
	}
)

var PaletteDefault = PaletteWhite

var PaletteMap = map[string]Palette{
	"mint":       PaletteMintGreen,
	"matrix":     PaletteMatrixGreen,
	"bog":        PaletteBogGreen,
	"swamp":      PaletteSwampGreen,
	"ice":        PaletteIceBlue,
	"desert":     PaletteDesert,
	"retro":      PaletteRetroIcyPole,
	"google":     PaletteGoogle,
	"default":    PaletteDefault,
	"white":      PaletteWhite,
	"black":      PaletteBlack,
	"red":        PaletteRed,
	"cyan":       PaletteCyan,
	"green":      PaletteGreen,
	"lime":       PaletteLime,
	"blue":       PaletteBlue,
	"yellow":     PaletteYellow,
	"purple":     PalettePurple,
	"red-black":  PaletteRedBlack,
	"red-orange": PaletteRedOrange,
	"pizza":      PalettePizza,
	"patriot":    PalettePatriot,
}

func GetPalette(p string) Palette {
	if val, ok := PaletteMap[p]; ok {
		return val
	}
	fmt.Println("Palette not found, using default palette")
	return GetPalette("default")
}

func GetColourMode(p string) ColourMode {
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
	case "l":
		fallthrough
	case "letter":
		return modeLetter
	case "patriot":
		return modePatriot
	}
	return modeNil
}
