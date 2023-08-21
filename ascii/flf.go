package ascii

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"io"
	"math"
	"math/rand"
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
	cMode := a.ColourMode
	if cMode == modeNil {
		cMode = a.Palette.ColourMode
	}
	var preRenderModes = []ColourMode{modeLetter}
	var postRenderModes = []ColourMode{modeSingle, modeAlternate, modeVerticalGradient, modeHorizontalGradient, modeStarsNStripes}
	letterList := f.getLetters(a.Message)

	if contains(preRenderModes, cMode) {
		switch cMode {
		case modeLetter:
			letterList = f.letterMode(a.Palette, letterList)
			break
		}

		renderedMsg := f.renderLetters(letterList)
		fmt.Println(renderedMsg)
	} else if contains(postRenderModes, cMode) {
		renderedMsg := f.renderLetters(letterList)

		switch cMode {
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
		case modeStarsNStripes:
			f.usaMode(renderedMsg)
			return
		}
	}
}

func (f *Font) getLetters(message string) [][]string {
	letterList := make([][]string, len(message))
	for i, char := range message {
		letter := f.charMap[char]
		cl := make([]string, len(letter))
		for r := 0; r < len(letter); r++ {
			cl[r] = strings.Clone(letter[r])
		}
		letterList[i] = cl
	}
	return letterList
}

func (f *Font) renderLetters(letterList [][]string) string {
	renderedMsg := ""
	for row := 0; row < f.height; row++ {
		for letter := 0; letter < len(letterList); letter++ {
			renderedMsg += letterList[letter][row]
		}
		renderedMsg += "\n"
	}

	var filteredLines string
	for _, l := range strings.Split(renderedMsg, "\n") {
		if strings.TrimSpace(l) != "" {
			filteredLines += l + "\n"
		}
	}

	return filteredLines
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

func (f *Font) usaMode(msg string) {
	lines := strings.Split(msg, "\n")
	renderStr := ""
	redLineIdx := -1

	flagStyle := color.S256(15, 17)
	//flagStyle := color.New(color.HEX("FFFFFF").Color(), color.HEX("0000ff", true).Color())
	//redLineStyle := color.New(color.FgBlack, color.BgRed)
	redLineStyle := color.HEX("ff0000")
	//whiteLineStyle := color.New(color.FgBlack, color.BgHiWhite)
	whiteLineStyle := color.HEX("FFFFFF")

	var filteredLines []string
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			filteredLines = append(filteredLines, l)
		}
	}

	for i, l := range filteredLines {
		if redLineIdx == -1 {
			redLineIdx = i
		}
		c := whiteLineStyle
		if i%2 == redLineIdx%2 {
			c = redLineStyle
		}

		if i <= len(filteredLines)/2 {
			chunks := cutIntoChunks(l, 3)
			flagPart := chunks[0]
			renderStr += flagStyle.Sprint(flagPart)

			rest := chunks[1] + chunks[2]
			renderStr += c.Sprint(rest)
		} else {
			renderStr += c.Sprint(l)
		}
		renderStr += "\n"
	}

	fmt.Println(renderStr)
}

func (f *Font) letterMode(p Palette, letters [][]string) [][]string {
	for letterIdx, letter := range letters {
		_, _ = color.Reset()
		colour := p.Colours[letterIdx%len(p.Colours)]
		for rowNum, rc := range letter {
			//fmt.Printf("lidx -> %d, %s %s\n", letterIdx, color.HEX(colour).Sprint(colour), color.HEX(colour).Sprint(rc))
			letters[letterIdx][rowNum] = color.HEX(colour).Sprint(rc)
		}
	}

	return letters
}

func contains(s []ColourMode, str ColourMode) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
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

func cutIntoChunks(l string, chunks int) []string {
	var result []string
	runes := []rune(l)

	cut := int(len(runes) / chunks)
	for i := 0; i < len(runes); i += cut {
		end := i + cut

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
