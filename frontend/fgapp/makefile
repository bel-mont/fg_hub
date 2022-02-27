# unix only
build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o bin/server

run: build
	./bin/server