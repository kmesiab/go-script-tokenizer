# Makefile for running Terraform and Go commands

# 🌍 Run app
run:
	@echo "Starting!"
	source .env && go build . && ./go-script-tokenizer.

# 🏗 Go build and test targets
build:
	@echo "🛠 Building Go project..."
	go build -o .

test:
	@echo "🚀 Running Go tests..."
	go test ./...

# 🌈 All-in-one linting
lint:
	@echo "🔍 Running all linters..."
	gofumpt -w . && golangci-lint run && markdownlint README.md

# 🌈 All-in-one build, test, and lint
all: build test lint
	@echo "🎉 Done!"

.PHONY: build test lint all
