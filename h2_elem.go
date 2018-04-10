// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// H2Elem is the React element definition corresponding to the HTML <h2> element
type H2Elem struct {
	Element
}

// _H2Props defines the properties for the <h2> element
type _H2Props struct {
	*BasicHTMLElement
}

// H2 creates a new instance of a <h2> element with the provided props and
// child
func H2(props *H2Props, children ...Element) *H2Elem {

	rProps := &_H2Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &H2Elem{
		Element: createElement("h2", rProps, children...),
	}
}
