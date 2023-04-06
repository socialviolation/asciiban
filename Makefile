BINARY="bin/asciiban"

build:
	go build -o $(BINARY) cli/main.go
	$(BINARY) "test message" -p retro
