package asciiban

import (
	"errors"
	"strconv"
	"strings"
)

type font struct {
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
func ParseFlf(fontName string, gz string) (*font, error) {
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

	f := &font{
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

func convertChar(font *font, char rune) ([]string, error) {
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
		revRow := Reverse(row)
		delim := revRow[0]
		if strings.HasSuffix(row, string(delim)+string(delim)) {
			row = Reverse(strings.Replace(revRow, string(delim), "", 2))
		} else if strings.HasSuffix(row, string(delim)) {
			row = Reverse(strings.Replace(revRow, string(delim), "", 1))
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

func (f *font) Render(word string) string {
	letterList := make([][]string, 0)
	for _, char := range word {
		letter := f.charMap[char]
		letterList = append(letterList, letter)
	}

	result := ""

	for row := 0; row < f.height; row++ {
		for letter := 0; letter < len(letterList); letter++ {
			result += letterList[letter][row]
		}
		result += "\n"
	}

	return result
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, mr := range s {
		n--
		runes[n] = mr
	}
	return string(runes[n:])
}
