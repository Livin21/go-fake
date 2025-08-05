# Makefile for go-fake

# Variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -s -w
BUILD_DIR = build

.PHONY: build test clean install help run-examples release release-local

# Build the application
build:
	mkdir -p bin
	go build -o bin/go-fake ./cmd/generate

# Build with version info
build-release:
	mkdir -p bin
	go build -ldflags="$(LDFLAGS)" -o bin/go-fake ./cmd/generate

# Release builds for all platforms
release:
	mkdir -p $(BUILD_DIR)
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-linux-amd64 ./cmd/generate
	# Linux ARM64  
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-linux-arm64 ./cmd/generate
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-windows-amd64.exe ./cmd/generate
	# Windows ARM64
	GOOS=windows GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-windows-arm64.exe ./cmd/generate
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-darwin-amd64 ./cmd/generate
	# macOS ARM64
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/go-fake-darwin-arm64 ./cmd/generate

# Create checksums for release builds
checksums: release
	cd $(BUILD_DIR) && sha256sum * > checksums.txt

# Local release build and test
release-local: checksums
	@echo "Release builds completed:"
	@ls -la $(BUILD_DIR)/
	@echo "\nChecksums:"
	@cat $(BUILD_DIR)/checksums.txt

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/ $(BUILD_DIR)/
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
	@echo "  build-release     - Build with optimizations for release"
	@echo "  release           - Build for all platforms"
	@echo "  checksums         - Create checksums for release builds"  
	@echo "  release-local     - Build and test release locally"
	@echo "  test              - Run tests"
	@echo "  clean             - Clean build artifacts"
	@echo "  deps              - Install dependencies"  
	@echo "  install           - Install binary to GOPATH/bin"
	@echo "  run-json          - Run with sample JSON schema"
	@echo "  run-sql           - Run with sample SQL schema"
	@echo "  run-comprehensive - Run with comprehensive example"
	@echo "  run-examples      - Run all examples"
	@echo "  help              - Show this help message"