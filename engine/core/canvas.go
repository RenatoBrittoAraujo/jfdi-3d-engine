package core

import (
	"syscall/js"
)

type Canvas struct {
	canvas js.Value
}

func CreateCanvas(context js.Value) *Canvas {
	return &Canvas{
		canvas: context,
	}
}

func (r *Canvas) moveTo(x int, y int) {
	r.canvas.Call("moveTo", x, y)
}

func (r *Canvas) lineTo(x int, y int) {
	r.canvas.Call("lineTo", x, y)
}

func (r *Canvas) stroke() {
	r.canvas.Call("stroke")
}
