package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello wasm!")
	document := js.Global().Get("document")
	target := document.Call("getElementById", "sample")
	button := document.Call("getElementById", "button")
	button.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		target.Set("innerHTML", "Hello, WebAssembry!!")
		return nil
	}))
	text := button.Get("innerText").String()
	fmt.Println(text)
	// プログラムが終了しないように待機
	// これをしないとClick発火時に Error: bad callback: Go program has already exited となる
	select {}
}
