package palettes

type Palette struct {
	// List of Hex Codes to use in scheme
	Palette []string
}

func (c *Palette) IsEmpty() bool {
	return c.Palette == nil
}

// https://colorhunt.co/palettes/neon
var (
	MintGreen    = Palette{Palette: []string{"CCFFBD", "CCFFBD", "7ECA9C", "7ECA9C", "4b795d", "4b795d"}}
	RetroIcyPole = Palette{Palette: []string{"F38181", "F38181", "FCE38A", "FCE38A", "EAFFD0", "95E1D3"}}
	Desert       = Palette{Palette: []string{"FFACAC", "FFBFA9", "FFBFA9", "FFEBB4", "FFEBB4", "FBFFB1"}}
	IceBlue      = Palette{Palette: []string{"E3FDFD", "E3FDFD", "CBF1F5", "A6E3E9", "A6E3E9", "71C9CE"}}
	SwampGreen   = Palette{Palette: []string{"DDFFBB", "DDFFBB", "C7E9B0", "C7E9B0", "B3C99C", "B3C99C"}}
	BogGreen     = Palette{Palette: []string{"7DCE13", "7DCE13", "70b911", "70b911", "64a40f", "64a40f"}}
	MatrixGreen  = Palette{Palette: []string{"00FF41", "00FF41", "00cc34", "00cc34", "009927", "009927"}}
	Google       = Palette{Palette: []string{"4285F4", "4285F4", "DB4437", "F4B400", "0F9D58"}}
	Default      = Palette{Palette: []string{"FFFFFF"}}
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
	"default": Default,
}

func Get(p string) Palette {
	if val, ok := ProfileMap[p]; ok {
		return val
	}
	return Default
}
