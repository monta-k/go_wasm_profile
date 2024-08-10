.PHONY: build
build:
	mkdir -p build
	tinygo build -o ./build/app.wasm -target wasm ./internal/main.go 