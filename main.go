package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"github.com/socialviolation/ascii-tester/fonts"
	"path"
	"strings"
)

func testAll(s string) {
	fonts := figure.AssetNames()
	for _, font := range fonts {

		fr := path.Base(font)
		fr = strings.Split(fr, ".")[0]

		fmt.Println()
		fmt.Println(fr)
		myFigure := figure.NewFigure(s, fr, false)
		myFigure.Print()
	}
}

func main() {
	raw := figure.NewFigureWithFont("Kasna: KIX", strings.NewReader(fonts.ANSIShadowStr), false).String()

	//https://colorhunt.co/palettes/neon
	//fgColors := []string{"CCFFBD", "CCFFBD", "7ECA9C", "7ECA9C", "4b795d", "4b795d"} // choc chip
	//fgColors := []string{"F38181", "F38181", "FCE38A", "FCE38A", "EAFFD0", "95E1D3"} // retro icy pole
	//fgColors := []string{"FFACAC", "FFBFA9", "FFBFA9", "FFEBB4", "FFEBB4", "FBFFB1"} // desert
	//fgColors := []string{"E3FDFD", "E3FDFD", "CBF1F5", "A6E3E9", "A6E3E9", "71C9CE"} // ice blue
	//fgColors := []string{"DDFFBB", "DDFFBB", "C7E9B0", "C7E9B0", "B3C99C", "B3C99C"} // swamp green
	//fgColors := []string{"7DCE13", "7DCE13", "70b911", "70b911", "64a40f", "64a40f"} // boogeyman green
	fgColors := []string{"00FF41", "00FF41", "00cc34", "00cc34", "009927", "009927"} // matrix green

	lines := strings.Split(raw, "\n")
	for i, l := range lines {
		if strings.Trim(l, " ") == "" {
			return
		}

		color.HEX(fgColors[i]).Println(l)
	}
}
