package main

import (
	"myitcv.io/react"

	"honnef.co/go/js/dom"
)

//go:generate reactGen

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app") // HL

	react.Render(App(), domTarget) // HL
}
