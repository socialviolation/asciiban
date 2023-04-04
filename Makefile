build:
	go build -o asciiban cmd/main.go

fonts:
	git clone git@github.com:xero/figlet-fonts.git

generate_fonts:
