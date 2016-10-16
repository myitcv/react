package main

import (
	r "github.com/myitcv/gopherjs/react"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	examplesDom := document.GetElementByID("examples")

	examples := Examples()

	r.Render(examples, examplesDom)
}
