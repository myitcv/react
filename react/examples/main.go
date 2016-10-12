package main

import (
	. "github.com/myitcv/gopherjs/react"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	examplesDom := document.GetElementByID("examples")

	examples := Examples()

	Render(examples, examplesDom)
}
