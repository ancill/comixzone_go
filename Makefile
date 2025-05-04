.PHONY: run build wasm clean run-wasm serve

# Development
run:
	air

# Build
build:
	go build -o bin/game ./cmd/game

# WASM build
wasm:
	GOOS=js GOARCH=wasm go build -o game.wasm ./cmd/game
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

# WASM hot reload
run-wasm:
	air -c .air.wasm.toml

# Serve files
serve:
	python3 -m http.server 8080

# Clean
clean:
	rm -rf bin/
	rm -f game.wasm wasm_exec.js

# Install air for hot reload
install-air:
	go install github.com/air-verse/air@latest

# Install dependencies
install:
	go mod tidy
	go mod download 