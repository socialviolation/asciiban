```text
 █████╗ ███████╗ ██████╗██╗██╗██████╗  █████╗ ███╗   ██╗
██╔══██╗██╔════╝██╔════╝██║██║██╔══██╗██╔══██╗████╗  ██║
███████║███████╗██║     ██║██║██████╔╝███████║██╔██╗ ██║
██╔══██║╚════██║██║     ██║██║██╔══██╗██╔══██║██║╚██╗██║
██║  ██║███████║╚██████╗██║██║██████╔╝██║  ██║██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝
```

asciiban is a command-line interface (CLI) tool written in Go that generates ASCII art banners for a given prompt. It is
a fun and creative way to make your command-line prompts stand out.

## Installation

You can install asciiban by running the following command:

```bash
go install github.com/social-violation/asciiban/cmd
```

This will install the tool in your `$GOBIN` directory.

## Usage

To use asciiban, simply run the following command:

```bash
asciiban "Your text here"
```

This will generate an ASCII art banner for the text you entered. You can customize the banner by using the following
flags:

* --font (-f): Specify the font to use for the banner. The ANSIShadow font is the default.
  * Available fonts can be found by running `asciiban list fonts`
* --palette (-p): Specify the colour palette to use.
  * Available fonts can be found by running `asciiban list palettes`

* For example, to generate an ASCII art banner using the big font and green color, run the following command:

```bash
asciiban -f big -c green "Your text here"
```

## Contributing

If you find any bugs or have suggestions for improvements, please feel free to open an issue or submit a pull request on
the GitHub repository: https://github.com/social-violation/asciiban

## License
This tool is released under the GPL-3.0 License. See the LICENSE file for details.
