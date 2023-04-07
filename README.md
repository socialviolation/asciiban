```text
 █████╗ ███████╗ ██████╗██╗██╗██████╗  █████╗ ███╗   ██╗
██╔══██╗██╔════╝██╔════╝██║██║██╔══██╗██╔══██╗████╗  ██║
███████║███████╗██║     ██║██║██████╔╝███████║██╔██╗ ██║
██╔══██║╚════██║██║     ██║██║██╔══██╗██╔══██║██║╚██╗██║
██║  ██║███████║╚██████╗██║██║██████╔╝██║  ██║██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝
```

asciiban is a go package or CLI tool that generates ASCII art banners for a given prompt. It is
a fun and creative way to make your command-line prompts stand out.

# SDK

# Installation
To Install, run the following tool.
`go get -u github.com/socialviolation/asciiban@main`

## Usage

It can then be used in your code using the following example:
```go
package main

import (
	"github.com/socialviolation/asciiban"
)

func main() {
    asciiban.Print(asciiban.Args{
        Message: "What is real?",
        Palette: palettes.MatrixGreen,
        Font:    fontpack.Georgi16,
    })
}

```

# CLI
## Installation

You can install asciiban by running the following command:

```bash
go install -o asciiban github.com/socialviolation/asciiban/cli@main
```

This will install the tool in your `$GOBIN` directory.

## Usage

To use asciiban, simply run the following command:

```bash
asciiban "Your text here"
```

This will generate an ASCII art banner for the text you entered. More information can be found by running `asciiban --help`

```text
Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fonts       A brief description of your command
  help        Help about any command
  palettes    A brief description of your command
  random      Generate Random ascii banner

Flags:
  -f, --font string      Colour palette to use (default "ansishadow")
  -h, --help             help for asciiban
  -m, --mode string      Palette Colour Mode (simple | alternating | vertical | horizontal)
  -p, --palette string   Colour palette to use (default "default")

```

* For example, to generate an ASCII art banner using the big font and green color, run the following command:

```bash
asciiban -f georgia11 -p matrix "What is real?"
```

## Contributing

If you find any bugs or have suggestions for improvements, please feel free to open an issue or submit a pull request on
the GitHub repository: https://github.com/socialviolation/asciiban

## License
This tool is released under the GPL-3.0 License. See the LICENSE file for details.
