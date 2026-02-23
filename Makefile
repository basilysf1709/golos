.PHONY: build install clean deps

BINARY=golos

build:
	go build -o $(BINARY) .

install: build
	sudo cp $(BINARY) /usr/local/bin/$(BINARY)

clean:
	rm -f $(BINARY)

deps:
	brew install portaudio
	go mod tidy

run: build
	./$(BINARY)
