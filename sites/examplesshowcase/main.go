// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("examples")

	examples := examples.Examples()

	r.Render(examples, domTarget)
}
