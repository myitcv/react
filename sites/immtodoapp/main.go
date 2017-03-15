package main

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples/immtodoapp"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("immtodoapp")

	examples := immtodoapp.TodoApp()

	r.Render(examples, domTarget)
}
