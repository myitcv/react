// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// MainElem is the React element definition corresponding to the HTML <main> element
type MainElem struct {
	Element
}

// _MainProps are the props for a <main> component
type _MainProps struct {
	*BasicHTMLElement
}

// Main creates a new instance of a <main> element with the provided props and children
func Main(props *MainProps, children ...Element) *MainElem {

	rProps := &_MainProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &MainElem{
		Element: createElement("main", rProps, children...),
	}
}
