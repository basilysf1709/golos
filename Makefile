.PHONY: build install clean deps run lint test

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

lint:
	@which golangci-lint > /dev/null 2>&1 || { echo "Installing golangci-lint..."; brew install golangci-lint; }
	golangci-lint run ./...

test:
	go test ./...
