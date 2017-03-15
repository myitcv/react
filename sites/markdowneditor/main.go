package main

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples/markdowneditor"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("markdowneditor")

	examples := markdowneditor.MarkdownEditor()

	r.Render(examples, domTarget)
}
