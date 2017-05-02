// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// H3Def is the React component definition corresponding to the HTML <h3> element
type H3Def struct {
	underlying *js.Object
}

// _H3Props defines the properties for the <h3> element
type _H3Props struct {
	*BasicHTMLElement
}

func (d *H3Def) reactElement() {}

// H3 creates a new instance of a <h3> element with the provided props and
// child
func H3(props *H3Props, children ...Element) *H3Def {

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

	return &H3Def{underlying: underlying}
}
