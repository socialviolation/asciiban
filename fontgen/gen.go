package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	dir, err := os.MkdirTemp("", "fontMap")
	if err != nil {
		log.Fatal(err)
	}
	defer func(path string) {
		fmt.Println("removing " + path)
		_ = os.RemoveAll(path)
	}(dir)

	repoUrl := "https://github.com/xero/figlet-fonts"
	_, _ = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	u := make([]string, 0)
	fontMap := make(map[string]string)
	for _, file := range files {
		fName := file.Name()
		if !strings.HasSuffix(fName, "flf") {
			continue
		}

		fmt.Println(fName)
		fName = strings.ReplaceAll(fName, " ", "")
		fName = strings.ReplaceAll(fName, "-", "")
		fName = strings.ReplaceAll(fName, "'", "")
		fName = strings.ReplaceAll(fName, "_", "")
		fName = strings.ReplaceAll(fName, ".flf", "")
		if fName[0] >= '0' && fName[0] <= '9' {
			fName = "F" + fName
		}
		uf := strings.ToLower(fName)
		if contains(u, uf) {
			continue
		}
		u = append(u, uf)

		fName = strings.Title(fName)
		b, _ := os.ReadFile(dir + "/" + file.Name())
		if !utf8.Valid(b) {
			continue
		}
		fc := string(b)
		fc = strings.ReplaceAll(fc, "`", "` + \"`\" + `")
		fontMap[fName] = fc
	}

	f, err := os.Create("fontpack/map.go")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	fontPack := template.Must(template.New("fontpack").Funcs(funcMap).Parse(fontPackTemplate))
	e := fontPack.ExecuteTemplate(f, "fontpack", TemplateArgs{
		Timestamp: time.Now(),
		URL:       repoUrl,
		FontMap:   fontMap,
	})

	if e != nil {
		fmt.Println("error rendering template: ", e)
		os.Exit(1)
	}

	for fontName, v := range fontMap {
		fontFile, _ := os.Create(fmt.Sprintf("fontpack/%s.go", strings.ToLower(fontName)))
		fontFileTemplate := template.Must(template.New("ff").Parse(specificFontTemplate))
		e = fontFileTemplate.ExecuteTemplate(fontFile, "ff", FontTemplateArgs{
			Timestamp:    time.Now(),
			URL:          repoUrl,
			FontName:     fontName,
			FontContents: v,
		})

		if e != nil {
			fmt.Println("error rendering template font template for : ", fontName, e)
			os.Exit(1)
		}

		_ = fontFile.Close()
	}

}

type TemplateArgs struct {
	Timestamp time.Time
	URL       string
	FontMap   map[string]string
}

type FontTemplateArgs struct {
	Timestamp    time.Time
	URL          string
	FontName     string
	FontContents string
}

var funcMap = template.FuncMap{
	"ToLower": strings.ToLower,
}

var fontPackTemplate = `// Package fontpack Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}
// using data from {{ .URL }}
package fontpack

import (
	"strings"
)

func Get(f string) string {
	if val, ok := FontMap[strings.ToLower(f)]; ok {
		return val
	}
	return ANSIShadow
}

var FontMap = map[string]string{
{{ range $key, $value := .FontMap }}	"{{ $key | ToLower}}": {{ $key }},
{{end }}}
`

var specificFontTemplate = `// Package fontpack Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}
// using data from {{ .URL }}
package fontpack

const {{ .FontName }} = ` + "`{{ .FontContents }}`" + `
`
