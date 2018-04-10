// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// H5Elem is the React element definition corresponding to the HTML <h5> element
type H5Elem struct {
	Element
}

// _H5Props defines the properties for the <h5> element
type _H5Props struct {
	*BasicHTMLElement
}

// H5 creates a new instance of a <h5> element with the provided props and
// child
func H5(props *H5Props, children ...Element) *H5Elem {

	rProps := &_H5Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &H5Elem{
		Element: createElement("h5", rProps, children...),
	}
}
