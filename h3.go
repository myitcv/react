// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// H3Elem is the React element definition corresponding to the HTML <h3> element
type H3Elem struct {
	Element
}

// _H3Props defines the properties for the <h3> element
type _H3Props struct {
	*BasicHTMLElement
}

// H3 creates a new instance of a <h3> element with the provided props and
// child
func H3(props *H3Props, children ...Element) *H3Elem {

	rProps := &_H3Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"h3", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &H3Elem{Element: elementHolder{elem: underlying}}
}
