// The following directive is necessary to make the package coherent:

//go:build fonts

package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

func eg() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
	fmt.Printf("  os.Args = %#v\n", os.Args)

	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println("  ", ev, "=", os.Getenv(ev))
	}
}

func main() {
	dir, err := os.MkdirTemp("", "fonts")
	if err != nil {
		log.Fatal(err)
	}
	defer func(path string) {
		fmt.Println("removing " + path)
		_ = os.RemoveAll(path)
	}(dir)

	_, _ = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/xero/figlet-fonts",
		Progress: os.Stdout,
	})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

const fontTemplate = "const {{.Name}} = \n`{{.Contents}}`\n"

const allFontsTemplate = "const AllFonts = []string{ {{.FontList}} }"
