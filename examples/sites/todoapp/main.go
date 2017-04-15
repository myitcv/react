// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	r "myitcv.io/react"
	"myitcv.io/react/examples/todoapp"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("todoapp")

	examples := todoapp.TodoApp()

	r.Render(examples, domTarget)
}
