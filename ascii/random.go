package ascii

import (
	"math/rand"
)

func allPalettes() []Palette {
	var a []Palette
	for _, v := range PaletteMap {
		a = append(a, v)
	}
	return a
}
func RandomPalette(p ...Palette) Palette {
	var pset []Palette
	if len(p) > 0 {
		pset = p
	} else {
		pset = allPalettes()
	}

	return pset[rand.Intn(len(pset))]
}

func allFonts() []string {
	var a []string
	for _, v := range fontMap {
		a = append(a, v)
	}
	return a
}
func RandomFont(p ...string) string {
	var fset []string
	if len(p) > 0 {
		fset = p
	} else {
		fset = allFonts()
	}

	return fset[rand.Intn(len(fset))]
}
