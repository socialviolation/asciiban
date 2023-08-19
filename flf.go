package asciiban

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"strconv"
	"strings"
)

type Font struct {
	fontName       string
	hardBlank      string
	height         int
	baseline       int
	maxLength      int
	oldLayout      int
	commentLines   int
	printDirection int
	fullLayout     int
	codetagCount   int

	startLine int
	fontSlice []string
	charMap   map[rune][]string
}

const minAscii = 32
const maxAscii = 127

/*
THE HEADER LINE

The header line gives information about the FIGfont.  Here is an example
showing the names of all parameters:

	        flf2a$ 6 5 20 15 3 0 143 229    NOTE: The first five characters in
	          |  | | | |  |  | |  |   |     the entire file must be "flf2a".
	         /  /  | | |  |  | |  |   \
	Signature  /  /  | |  |  | |   \   Codetag_Count
	  Hardblank  /  /  |  |  |  \   Full_Layout*
	       Height  /   |  |   \  Print_Direction
	       Baseline   /    \   Comment_Lines
	        Max_Length      Old_Layout*

	* The two layout parameters are closely related and fairly complex.
	    (See "INTERPRETATION OF LAYOUT PARAMETERS".)
*/

// ParseFlf parses a FIGlet font file
func ParseFlf(fontName string, gz string) (*Font, error) {
	cont, err := readCompressedFont(gz)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(cont, "\n")
	if len(lines) < 1 {
		return nil, errors.New("font content error")
	}
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Replace(lines[i], "\r", "", 1)
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	header := strings.Split(lines[0], " ")
	oldHeader := len(header) < 9

	f := &Font{
		fontName:  fontName,
		hardBlank: header[0][len(header[0])-1:],
		charMap:   make(map[rune][]string),
	}

	f.height, _ = strconv.Atoi(header[1])
	f.startLine = len(lines) - (f.height * (maxAscii - minAscii))
	f.baseline, _ = strconv.Atoi(header[2])
	f.maxLength, _ = strconv.Atoi(header[3])
	f.oldLayout, _ = strconv.Atoi(header[4])
	f.commentLines, _ = strconv.Atoi(header[5])
	if !oldHeader {
		f.printDirection, _ = strconv.Atoi(header[6])
		f.fullLayout, _ = strconv.Atoi(header[7])
		f.codetagCount, _ = strconv.Atoi(header[8])
	}

	f.fontSlice = lines[f.commentLines+1:]

	for _, asciiChar := range makeRange(minAscii, maxAscii) {
		f.charMap[rune(asciiChar)], _ = convertChar(f, rune(asciiChar))
	}

	return f, nil
}

func convertChar(font *Font, char rune) ([]string, error) {
	if char < minAscii || char > maxAscii {
		return nil, errors.New("not Ascii")
	}

	beginRow := (int(char) - minAscii) * font.height
	word := make([]string, font.height)

	for i := 0; i < font.height; i++ {
		idx := beginRow + i
		if idx >= len(font.fontSlice) {
			continue
		}
		row := font.fontSlice[idx]
		revRow := reverse(row)
		delim := revRow[0]
		if strings.HasSuffix(row, string(delim)+string(delim)) {
			row = reverse(strings.Replace(revRow, string(delim), "", 2))
		} else if strings.HasSuffix(row, string(delim)) {
			row = reverse(strings.Replace(revRow, string(delim), "", 1))
		}

		row = strings.Replace(row, font.hardBlank, " ", -1)
		word[i] = row
	}

	return word, nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (f *Font) Render(a Args) {
	letterList := make([][]string, 0)
	for _, char := range a.Message {
		letter := f.charMap[char]
		letterList = append(letterList, letter)
	}

	var preRenderModes = []ColourMode{modeLetter}
	var postRenderModes = []ColourMode{modeSingle, modeAlternate, modeVerticalGradient, modeHorizontalGradient}

	if contains(preRenderModes, a.ColourMode) {
		fmt.Println("TODO: NOT IMPLEMENTED")
	} else if contains(postRenderModes, a.ColourMode) {
		renderedMsg := ""
		for row := 0; row < f.height; row++ {
			for letter := 0; letter < len(letterList); letter++ {
				renderedMsg += letterList[letter][row]
			}
			renderedMsg += "\n"
		}
		switch a.ColourMode {
		case modeSingle:
			f.singleColour(a.Palette, renderedMsg)
			return
		case modeAlternate:
			f.alternatingColours(a.Palette, renderedMsg)
			return
		case modeVerticalGradient:
			f.verticalGradient(a.Palette, renderedMsg)
			return
		case modeHorizontalGradient:
			f.horizontalGradient(a.Palette, renderedMsg)
			return
		}
	}
}

func (f *Font) singleColour(p Palette, msg string) {
	color.HEX(p.Colours[0]).Println(msg)
}

func (f *Font) alternatingColours(p Palette, msg string) {
	lines := strings.Split(msg, "\n")
	for i, l := range lines {
		n := i % len(p.Colours)
		if n >= len(p.Colours) {
			n = 0
		}
		color.HEX(p.Colours[n]).Println(l)
	}
}

func (f *Font) verticalGradient(p Palette, msg string) {
	lines := strings.Split(msg, "\n")
	palLen := len(p.Colours)
	for i, l := range lines {
		ind := translateLERP(len(lines), palLen, i)
		color.HEX(p.Colours[ind]).Println(l)
	}
}

func (f *Font) horizontalGradient(p Palette, msg string) {
	lines := strings.Split(msg, "\n")
	longest := getLongestString(lines)
	chunkSize := (longest / len(p.Colours)) + 1
	for _, l := range lines {
		if l == "" {
			continue
		}
		lineChunks := sliceIntoChunks(l, chunkSize)
		for c := 0; c < len(lineChunks); c++ {
			color.HEX(p.Colours[c]).Print(lineChunks[c])
		}
		fmt.Println()
	}
}

func contains(s []ColourMode, str ColourMode) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
