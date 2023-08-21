# ASCIIBAN CLI

## Installation

You can install asciiban cli by running the following command:

```bash
go install github.com/socialviolation/asciiban/asciicli@main
```

This will install the tool in your `$GOBIN` directory.

## Usage

To use asciiban cli, simply run the following command:

```bash
asciicli "Your text here"
```

This will generate an ASCII art banner for the text you entered. More information can be found by running `asciicli --help`

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