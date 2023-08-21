set dotenv-load
set fallback := false

BINDIR := justfile_directory() + '/bin'
BINARY := "asciicli"

@default:
	just --list --unsorted

build:
	cd asciicli && go build -o {{BINDIR}}/{{BINARY}} main.go
	{{BINDIR}}/{{BINARY}} "test message" -p retro

generate:
    go generate ./...
