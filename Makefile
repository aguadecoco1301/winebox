.PHONY: build run install uninstall clean release

BINARY=winebox
PREFIX=/usr/local

build:
	go build -o $(BINARY)

run:
	go run . $(ARGS)

install: build
	install -Dm755 $(BINARY) $(PREFIX)/bin/$(BINARY)

uninstall: 
	rm -rf $(PREFIX)/bin/$(BINARY)

clean:
	rm -rf $(BINARY)

release:
	mkdir -p dist
	go build -ldflags="-s -w" -o dist/$(BINARY)

