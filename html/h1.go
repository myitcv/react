// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// H1Elem is the React element definition corresponding to the HTML <h1> element
type H1Elem struct {
	react.Element
}

// _H1Props defines the properties for the <h1> element
type _H1Props struct {
	*BasicHTMLElement
}

// H1 creates a new instance of a <h1> element with the provided props and
// child
func H1(props *H1Props, children ...react.Element) *H1Elem {

	rProps := &_H1Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &H1Elem{
		Element: react.InternalCreateElement("h1", rProps, children...),
	}
}