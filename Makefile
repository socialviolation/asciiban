build:
	go build -o bin/asciiban cmd/main.go
	asciiban "test message" -p retro

fonts:
	git clone git@github.com:xero/figlet-fonts.git
