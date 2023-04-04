package asciiban

type ColourProfile struct {
	// List of Hex Codes to use in scheme
	Palette []string
}

func (c *ColourProfile) isEmpty() bool {
	return c.Palette == nil
}

// https://colorhunt.co/palettes/neon
var (
	MintGreen    = ColourProfile{Palette: []string{"CCFFBD", "CCFFBD", "7ECA9C", "7ECA9C", "4b795d", "4b795d"}}
	RetroIcyPole = ColourProfile{Palette: []string{"F38181", "F38181", "FCE38A", "FCE38A", "EAFFD0", "95E1D3"}}
	Desert       = ColourProfile{Palette: []string{"FFACAC", "FFBFA9", "FFBFA9", "FFEBB4", "FFEBB4", "FBFFB1"}}
	IceBlue      = ColourProfile{Palette: []string{"E3FDFD", "E3FDFD", "CBF1F5", "A6E3E9", "A6E3E9", "71C9CE"}}
	SwampGreen   = ColourProfile{Palette: []string{"DDFFBB", "DDFFBB", "C7E9B0", "C7E9B0", "B3C99C", "B3C99C"}}
	BogGreen     = ColourProfile{Palette: []string{"7DCE13", "7DCE13", "70b911", "70b911", "64a40f", "64a40f"}}
	MatrixGreen  = ColourProfile{Palette: []string{"00FF41", "00FF41", "00cc34", "00cc34", "009927", "009927"}}
	Default      = ColourProfile{Palette: []string{"FFFFFF"}}
)

var Profiles = map[string]ColourProfile{
	"mint":    MintGreen,
	"matrix":  MatrixGreen,
	"bog":     BogGreen,
	"swamp":   SwampGreen,
	"ice":     IceBlue,
	"desert":  Desert,
	"retro":   RetroIcyPole,
	"default": Default,
}

func GetProfile(p string) ColourProfile {
	if val, ok := Profiles[p]; ok {
		return val
	}
	return Default
}
