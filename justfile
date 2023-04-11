set dotenv-load
set fallback := false

BINDIR := justfile_directory() + '/bin'
BINARY := "asciiban"

@default:
	just --list --unsorted

build:
	cd cli && go work sync
	cd cli && go build -o {{BINDIR}}/{{BINARY}} main.go
	{{BINARY}} "test message" -p retro
