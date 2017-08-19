// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"net/url"

	"myitcv.io/react"
	"myitcv.io/react/examples/hellomessage"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	domTarget := document.GetElementByID("hellomessage")

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
				Href:      "https://github.com/myitcv/react/blob/master/examples/sites/timer/main.go",
				Title:     "Source on GitHub",
			},
			react.S("Source on GitHub"),
		)

		elems = append(elems, a)
	}

	props := hellomessage.HelloMessageProps{
		Name: "Jane",
	}

	elems = append(elems, hellomessage.HelloMessage(props))

	react.Render(react.Div(nil, elems...), domTarget)
}
