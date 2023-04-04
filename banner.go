package asciiban

import (
	"fmt"
	"path"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

func GetFonts() {
	for _, font := range figure.AssetNames() {

		fr := path.Base(font)
		fr = strings.Split(fr, ".")[0]

		fmt.Println()
		fmt.Println(fr)
		myFigure := figure.NewFigure("asciiban", fr, false)
		myFigure.Print()
	}
}

type Args struct {
	Message string
	Font    string
	Profile ColourProfile
}

func Print(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), true).String()

	palLen := len(args.Profile.Palette)
	lines := strings.Split(raw, "\n")
	for i, l := range lines {
		if strings.Trim(l, " ") == "" {
			continue
		}
		ind := i
		if i >= palLen {
			ind = palLen - 1
		}

		color.HEX(args.Profile.Palette[ind]).Println(l)
	}
}
