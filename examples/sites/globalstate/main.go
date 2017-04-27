// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	r "myitcv.io/react"
	"myitcv.io/react/examples/sites/globalstate/model"
	"myitcv.io/react/examples/sites/globalstate/state"

	"honnef.co/go/js/dom"
)

//go:generate reactGen

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app")

	state.State.Root().People().Set(model.NewPeople(
		model.NewPerson("Peter", 50),
		model.NewPerson("Paul", 51),
		model.NewPerson("Mary", 52),
	))

	r.Render(App(), domTarget)
}
