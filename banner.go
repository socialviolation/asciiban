//go:build !exclude

//go:generate go run gen.go
package asciiban

import (
	"bytes"
	"compress/gzip"
	"io"
	"math"
	"math/rand"
	"strings"
)

type Args struct {
	Message    string
	Font       *Font
	Palette    Palette
	ColourMode ColourMode
}

var DefaultArgs Args = Args{
	Message: "ascii banner",
	Font:    nil,
	Palette: PaletteDefault,
}

func Print(args Args) {
	if args.Font == nil {
		args.Font, _ = GetFont("default")
	}
	if args.Palette.IsEmpty() {
		args.Palette = White
	}
	if args.ColourMode != modeNil {
		args.Palette.ColourMode = args.ColourMode
	}

	args.Font.Render(args)
}

func Random(args Args) {
	var err error
	args.Font, err = GetFont(pick(FontMap))
	if err != nil {
		panic(err)
	}
	args.Palette = pick(ProfileMap)
	args.ColourMode = modeNil

	args.Font.Render(args)
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

func readCompressedFont(filePath string) (string, error) {
	reader := strings.NewReader(filePath)

	// Create a new gzip reader
	gzipReader, err := gzip.NewReader(reader)
	defer gzipReader.Close()

	buffer := bytes.NewBufferString("")
	// Copy the contents of the gzip reader to the buffer string
	_, err = io.Copy(buffer, gzipReader)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, mr := range s {
		n--
		runes[n] = mr
	}
	return string(runes[n:])
}
