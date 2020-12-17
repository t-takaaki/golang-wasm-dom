# Go Wasm Dom
Go 1.15
You can get started with:
```
$ GOOS=js GOARCH=wasm go build -o ./public/sample.wasm sample.go
$ node server.js 
```
Then browse to the running server: http://localhost:3000/

- https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
- https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html
