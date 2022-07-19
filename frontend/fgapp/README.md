# FG App

WASM Application

## Building and running

Windows compilation might fail. If so, attempt WSL2.

### If makefile is available

`make run`

Otherwise, build both the client and the server separately.

### Client

`GOOS=js GOARCH=wasm go build -o ../../web/assets/wasm.wasm`

### Server

`go build -o bin/server.exe`

### Run

`./bin/server.exe`
