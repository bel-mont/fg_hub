# FG App
Built with [Go App](https://github.com/maxence-charriere/go-app).


## Building and running
### If makefile is available
`make run`

Otherwise, build both the client and the server separately.
### Client
`GOARCH=wasm GOOS=js go build -o web/app.wasm`
### Server
`go build -o bin/server.exe` 

### Run
`./bin/server.exe`
