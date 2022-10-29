set GOOS=js
set GOARCH=wasm
go build -o myWasm.wasm myWasm.go canvas2d.go