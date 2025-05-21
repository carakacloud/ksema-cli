.PHONY: build run clean

build:
	CGO_ENABLED=0 go build -o bin/ksema-cli ./cmd/ksema-cli

run: build
	./bin/ksema-cli

clean:
	rm -rf bin/ksema-cli
