// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"myitcv.io/react"
	"myitcv.io/react/examples/immtodoapp"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("immtodoapp")

	examples := immtodoapp.TodoApp()

	react.Render(examples, domTarget)
}
