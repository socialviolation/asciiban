package asciiban

import (
	"github.com/socialviolation/asciiban/cprofiles"
	"github.com/socialviolation/asciiban/fonts"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

const backgroundChar = "░"

type Args struct {
	Message string
	Font    string
	Profile cprofiles.ColourProfile
	FillBg  bool
}

var DefaultArgs Args = Args{
	Message: "asciiban",
	Font:    fonts.ANSIShadow,
	Profile: cprofiles.Default,
	FillBg:  false,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = fonts.ANSIShadow
	}
	if args.Profile.IsEmpty() {
		args.Profile = cprofiles.Default
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
