// main.go
package main

import (
	"fmt"
	"syscall/js"

	"github.com/nanlei2000/pi/pi"
)

func calc(this js.Value, args []js.Value) any {
	n := args[0].Int()
	piRes := pi.Calc(uint64(n))

	fmt.Println(piRes)

	return 1
}

func main() {
	// 将golang中calc函数注入到window.calc中
	js.Global().Set("calc", js.FuncOf(calc))

	select {}
}
