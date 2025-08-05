# Makefile for go-fake

.PHONY: build test clean install help run-examples

# Build the application
build:
	mkdir -p bin
	go build -o bin/go-fake ./cmd/generate

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install the binary to GOPATH/bin
install: build
	cp bin/go-fake $(GOPATH)/bin/

# Run with sample JSON
run-json: build
	./bin/go-fake -schema examples/sample.json -output sample-output.csv

# Run with sample SQL
run-sql: build
	./bin/go-fake -schema examples/simple.sql -output sql-output.csv

# Run comprehensive example
run-comprehensive: build
	./bin/go-fake -schema examples/comprehensive.json -output comprehensive-output.csv -rows 50

# Run all examples
run-examples: run-json run-sql run-comprehensive

# Show help
help:
	@echo "Available commands:"
	@echo "  build             - Build the application"
	@echo "  test              - Run tests"
	@echo "  clean             - Clean build artifacts"
	@echo "  deps              - Install dependencies"  
	@echo "  install           - Install binary to GOPATH/bin"
	@echo "  run-json          - Run with sample JSON schema"
	@echo "  run-sql           - Run with sample SQL schema"
	@echo "  run-comprehensive - Run with comprehensive example"
	@echo "  run-examples      - Run all examples"
	@echo "  help              - Show this help message"