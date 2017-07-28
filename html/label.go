// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// LabelElem is the React element definition corresponding to the HTML <label> element
type LabelElem struct {
	react.Element
}

// _LabelProps defines the properties for the <label> element
type _LabelProps struct {
	*BasicHTMLElement

	For string `js:"htmlFor"`
}

// Label creates a new instance of a <label> element with the provided props and child
// element
func Label(props *LabelProps, child react.Element) *LabelElem {

	rProps := &_LabelProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &LabelElem{
		Element: react.InternalCreateElement("label", rProps, child),
	}
}
