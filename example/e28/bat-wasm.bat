set GOARCH=wasm
set GOOS=js
go build -o main.wasm main.go
