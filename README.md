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
        Palette: asciiban.PaletteMintGreen,
        Font:    fontpack.Georgi16,
    })
}
```

# CLI
## Installation

You can install the asciiban cli by running the following command:

```bash
go install github.com/socialviolation/asciiban/asciicli@main
```

This will install the tool in your `$GOBIN` directory.

## Usage

To use asciiban cli, simply run the following command:

```bash
asciicli "Your text here"
```

This will generate an ASCII art banner for the text you entered. More information can be found by running `asciiban --help`

```text
Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fonts       Subcommands show info for available fonts
  help        Help about any command
  palettes    Subcommands show info for available palettes
  random      Generate ascii banner using random font & colours

Flags:
  -f, --font string      Colour palette to use (default "ansishadow")
  -h, --help             help for asciicli
  -m, --mode string      Palette Colour Mode (simple | alternating | vertical | horizontal)
  -p, --palette string   Colour palette to use (default "default")

```

* For example, to generate an ASCII art banner using the big font and green color, run the following command:

```bash
asciicli -f georgia11 -p matrix "What is real?"
```

## Contributing

If you find any bugs or have suggestions for improvements, please feel free to open an issue or submit a pull request on
the GitHub repository: https://github.com/socialviolation/asciiban

## License
This tool is released under the GPL-3.0 License. See the LICENSE file for details.

```shell
$ asciicli "Good Yard" -f caligraphy2

               .-'''-.         .-'''-.
              '   _    \      '   _    \   _______                                              _______
            /   /` '.   \   /   /` '.   \  \  ___ `'.                                           \  ___ `'.
  .--./)   .   |     \  '  .   |     \  '   ' |--.\  \      .-.          .-                      ' |--.\  \
 /.''\\    |   '      |  ' |   '      |  '  | |    \  '      \ \        / /            .-,.--.   | |    \  '
| |  | |   \    \     / /  \    \     / /   | |     |  '      \ \      / /      __     |  .-. |  | |     |  '
 \`-' /     `.   ` ..' /    `.   ` ..' /    | |     |  |       \ \    / /    .:--.'.   | |  | |  | |     |  |
 /("'`         '-...-'`        '-...-'`     | |     ' .'        \ \  / /    / |   \ |  | |  | |  | |     ' .'
 \ '---.                                    | |___.' /'          \ `  /     `" __ | |  | |  '-   | |___.' /'
  /'""'.\                                  /_______.'/            \  /       .'.''| |  | |      /_______.'/
 ||     ||                                 \_______|/             / /       / /   | |_ | |      \_______|/
 \'. __//                                                     |`-' /        \ \._,\ '/ |_|
  `'---'                                                       '..'          `--'  `"
```
