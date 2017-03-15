package main

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples/hellomessage"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("hellomessage")

	props := hellomessage.HelloMessageProps{
		Name: "Jane",
	}

	examples := hellomessage.HelloMessage(props)

	r.Render(examples, domTarget)
}
