set dotenv-load
set fallback := false

BINDIR := justfile_directory() + '/bin'
BINARY := "asciiban"

@default:
	just --list --unsorted

build:
	go build -o {{BINDIR}}/{{BINARY}} main.go
	{{BINDIR}}/{{BINARY}} "test message" -p retro

generate:
    go generate ./...

output_readme_screens:
    {{BINARY}} "america" -p patriot -f dosrebel
    {{BINARY}} "what is real?" -p matrix -f georgia11
    {{BINARY}} "Google" -p google -f univers
    {{BINARY}} "Good yard" -f crazy -p retro

