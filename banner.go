package asciiban

import (
	"github.com/socialviolation/asciiban/fonts"
	"github.com/socialviolation/asciiban/palettes"
	"log"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

const backgroundChar = "░"

type Args struct {
	Message string
	Font    string
	Profile palettes.Palette
	FillBg  bool
}

var DefaultArgs Args = Args{
	Message: "asciiban",
	Font:    fonts.ANSIShadow,
	Profile: palettes.Default,
	FillBg:  false,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = fonts.ANSIShadow
	}
	if args.Profile.IsEmpty() {
		args.Profile = palettes.Default
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
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
