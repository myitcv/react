// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"myitcv.io/react"
)

// IElem is the React element definition corresponding to the HTML <i> element
type IElem struct {
	react.Element
}

// _IProps are the props for a <i> component
type _IProps struct {
	*BasicHTMLElement

	Src string `js:"src"`
}

// I creates a new instance of a <i> element with the provided props and children
func I(props *IProps, children ...react.Element) *IElem {

	rProps := &_IProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &IElem{
		Element: react.InternalCreateElement("i", rProps, children...),
	}
}