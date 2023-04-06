//go:generate go run fontgen/gen.go
package asciiban

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"github.com/socialviolation/asciiban/fontpack"
	"github.com/socialviolation/asciiban/palettes"
	"log"
	"math"
	"strings"
)

type Args struct {
	Message    string
	Font       string
	Palette    palettes.Palette
	ColourMode palettes.ColourMode
}

var DefaultArgs Args = Args{
	Message: "asciiban",
	Font:    fontpack.ANSIShadow,
	Palette: palettes.White,
}

func Print(args Args) {
	if args.Font == "" {
		args.Font = fontpack.ANSIShadow
	}
	if args.Palette.IsEmpty() {
		args.Palette = palettes.White
	}
	if args.ColourMode != palettes.Nil {
		args.Palette.ColourMode = args.ColourMode
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	switch args.Palette.ColourMode {
	case palettes.Single:
		printSingleColour(args)
		return
	case palettes.Alternate:
		printAlternatingColours(args)
		return
	case palettes.VerticalGradient:
		printVerticalGradient(args)
		return
	case palettes.HorizontalGradient:
		printHorizontalGradient(args)
		return
	}
}

func printSingleColour(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	color.HEX(args.Palette.Colours[0]).Println(raw)
}

func printAlternatingColours(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	raw = strings.TrimSuffix(raw, " ")
	lines := strings.Split(raw, "\n")
	for i, l := range lines {
		if i == len(lines)-1 {
		}

		if strings.Trim(l, " ") == "" {
			continue
		}

		n := i % 2
		if n >= len(args.Palette.Colours) {
			n = 0
		}
		color.HEX(args.Palette.Colours[n]).Println(l)
	}
}

func printVerticalGradient(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	raw = strings.TrimSuffix(raw, " ")
	lines := strings.Split(raw, "\n")
	palLen := len(args.Palette.Colours)
	for i, l := range lines {
		if strings.Trim(l, " ") == "" {
			continue
		}
		ind := translateLERP(len(lines), palLen, i)
		color.HEX(args.Palette.Colours[ind]).Println(l)
	}
}

func printHorizontalGradient(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	lines := strings.Split(raw, "\n")
	palLen := len(args.Palette.Colours)
	for _, l := range lines {
		if strings.Trim(l, " ") == "" {
			continue
		}

		lineChunks := chunkSlice(l, palLen)
		for c := 0; c < len(lineChunks); c++ {
			color.HEX(args.Palette.Colours[c]).Print(lineChunks[c])
		}
		fmt.Println()
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

func chunkSlice(slice string, numChunks int) []string {
	var result []string

	for i := 0; i < numChunks; i++ {

		min := i * len(slice) / numChunks
		max := ((i + 1) * len(slice)) / numChunks

		result = append(result, slice[min:max])
	}
	return result

}
