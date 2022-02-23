# unix only
build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o bin/server.exe

run: build
	./bin/server.exe