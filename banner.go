//go:generate go run gen.go
package asciiban

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"github.com/socialviolation/asciiban/fontpack"
	"github.com/socialviolation/asciiban/palettes"
	"log"
	"math"
	"math/rand"
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

func Random(args Args) {
	args.Font = pick(fontpack.FontMap)
	args.Palette = pick(palettes.ProfileMap)
	args.ColourMode = palettes.Nil
	Print(args)
}

func printSingleColour(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	color.HEX(args.Palette.Colours[0]).Println(raw)
}

func printAlternatingColours(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	lines := strings.Split(raw, "\n")
	for i, l := range lines {
		n := i % 2
		if n >= len(args.Palette.Colours) {
			n = 0
		}
		color.HEX(args.Palette.Colours[n]).Println(l)
	}
}

func printVerticalGradient(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	lines := strings.Split(raw, "\n")
	palLen := len(args.Palette.Colours)
	for i, l := range lines {
		ind := translateLERP(len(lines), palLen, i)
		color.HEX(args.Palette.Colours[ind]).Println(l)
	}
}

func printHorizontalGradient(args Args) {
	raw := figure.NewFigureWithFont(args.Message, strings.NewReader(args.Font), false).String()
	lines := strings.Split(raw, "\n")
	longest := getLongestString(lines)
	chunkSize := (longest / len(args.Palette.Colours)) + 1
	for _, l := range lines {
		if l == "" {
			continue
		}
		lineChunks := sliceIntoChunks(l, chunkSize)
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

func pick[K comparable, V any](m map[K]V) V {
	k := rand.Intn(len(m))
	i := 0
	for _, x := range m {
		if i == k {
			return x
		}
		i++
	}
	panic("unreachable")
}

func getLongestString(slice []string) int {
	longest := 0
	for _, s := range slice {
		r := []rune(s)
		if len(r) > longest {
			longest = len(r)
		}
	}
	return longest
}

func sliceIntoChunks(l string, chunkSize int) []string {
	var result []string
	runes := []rune(l)
	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize

		if end > len(runes) {
			end = len(runes)
		}

		result = append(result, string(runes[i:end]))
	}

	return result
}
