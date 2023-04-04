build:
	go build -o asciiban cmd/main.go
	./asciiban print -m "asciiban"

fonts:
	git clone git@github.com:xero/figlet-fonts.git
