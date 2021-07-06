package core

import (
	"fmt"
	"syscall/js"
)

func SetupEngine() {
	fmt.Println("SETUP ENGINE START")

	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "id")
	context := canvas.Call("getContext", "2d")

	newCanvas := CreateCanvas(context)
	newCanvas.moveTo(0, 0)
	newCanvas.lineTo(200, 100)
	newCanvas.stroke()

	fmt.Println("SETUP ENGINE END")
}
