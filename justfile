set dotenv-load
set fallback := false

BINDIR := justfile_directory() + '/bin'
BINARY := "cli"

@default:
	just --list --unsorted

build:
	cd cli && go build -o {{BINDIR}}/{{BINARY}} main.go
	{{BINARY}} "test message" -p retro
