package asciiban

import (
	"github.com/socialviolation/asciiban/fonts"
	"github.com/socialviolation/asciiban/palettes"
	"log"
	"math"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

const backgroundChar = "░"

type Args struct {
	Message string
	Font    string
	Palette palettes.Palette
	FillBg  bool
}

var DefaultArgs Args = Args{
	Message: "asciiban",
	Font:    fonts.ANSIShadow,
	Palette: palettes.White,
	FillBg:  false,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = fonts.ANSIShadow
	}
	if args.Palette.IsEmpty() {
		args.Palette = palettes.White
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

	raw = strings.TrimSuffix(raw, " ")
	lines := strings.Split(raw, "\n")
	palLen := len(args.Palette.Colours)
	for i, l := range lines {
		if i == len(lines)-1 {
		}

		if strings.Trim(l, " ") == "" {
			continue
		}
		ind := translateLERP(len(lines), palLen, i)
		color.HEX(args.Palette.Colours[ind]).Println(l)
	}
}

func translateLERP(lines int, colours int, lineIndex int) int {
	transInd := float64(lineIndex) / float64(lines)
	ci := lerp(0, colours, transInd)
	if ci >= colours {
		ci = colours - 1
	}
	return ci
}

func lerp(x int, y int, f float64) int {
	i := float64(x) + f*(float64(y)-float64(x))
	return int(math.Round(i))
}
