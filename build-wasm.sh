#!/bin/bash
GOOS=js GOARCH=wasm go build -o game.wasm ./cmd/game
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
echo "WASM build completed" 