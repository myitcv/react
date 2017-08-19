// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"net/url"

	"myitcv.io/react"
	"myitcv.io/react/examples/sites/globalstate/model"
	"myitcv.io/react/examples/sites/globalstate/state"

	"honnef.co/go/js/dom"
)

//go:generate reactGen

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	domTarget := document.GetElementByID("app")

	u, err := url.Parse(document.URL())
	if err != nil {
		panic(err)
	}

	var elems []react.Element

	if u.Query().Get("hideGithubRibbon") != "true" {
		a := react.A(
			&react.AProps{
				ClassName: "github-fork-ribbon right-top",
				Target:    "_blank",
				Href:      "https://github.com/myitcv/react/blob/master/examples/sites/globalstate/main.go",
				Title:     "Source on GitHub",
			},
			react.S("Source on GitHub"),
		)

		elems = append(elems, a)
	}

	elems = append(elems, App())

	state.State.Root().People().Set(model.NewPeople(
		model.NewPerson("Peter", 50),
		model.NewPerson("Paul", 51),
		model.NewPerson("Mary", 52),
	))

	react.Render(react.Div(nil, elems...), domTarget)
}
