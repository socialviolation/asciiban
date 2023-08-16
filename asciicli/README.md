# ASCIIBAN CLI

## Installation

You can install asciiban cli by running the following command:

```bash
go install github.com/socialviolation/asciiban/asciicli-bkp@main
```

This will install the tool in your `$GOBIN` directory.

## Usage

To use asciiban cli, simply run the following command:

```bash
asciicli-bkp "Your text here"
```

This will generate an ASCII art banner for the text you entered. More information can be found by running `asciicli --help`

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
asciicli-bkp -f georgia11 -p matrix "What is real?"
```