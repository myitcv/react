// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// UlElem is the React element definition corresponding to the HTML <ul> element
type UlElem struct {
	react.Element
}

// _UlProps defines the properties for the <ul> element
type _UlProps struct {
	*BasicHTMLElement
}

// Ul creates a new instance of a <ul> element with the provided props and <li>
// children
func Ul(props *UlProps, children ...*LiElem) *UlElem {

	rProps := &_UlProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	var elems []react.Element
	for _, v := range children {
		elems = append(elems, v)
	}

	return &UlElem{
		Element: react.InternalCreateElement("ul", rProps, elems...),
	}
}