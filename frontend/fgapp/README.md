# FG App
WASM Application

u
## Building and running
Windows compilation might fail. If so, attempt WSL2.
### If makefile is available
`make run`

Otherwise, build both the client and the server separately.
### Client
Use TinyGo WASM, to keep file size down.
`tinygo build -o ../../web/assets/wasm.wasm -target wasm ./main.go`
Make sure that wasm_exec is also available.
`cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js wasm_exec_tiny.js`

`GOOS=js GOARCH=wasm go build -o  ../../web/assets/wasm.wasm`
### Server
`go build -o bin/server.exe` 

### Run
`./bin/server.exe`
