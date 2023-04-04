package asciiban

import (
	"fmt"
	"github.com/socialviolation/asciiban/fonts"
	"path"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

const backgroundChar = "░"

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
	FillBg  bool
}

var DefaultArgs Args = Args{
	Message: "asciiban",
	Font:    fonts.ANSIShadow,
	Profile: Default,
	FillBg:  false,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = fonts.ANSIShadow
	}
	if args.Profile.isEmpty() {
		args.Profile = Default
	}

	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), true).String()
	if args.FillBg {
		raw = strings.Replace(raw, " ", backgroundChar, -1)
	}

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
