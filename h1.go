// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// H1Def is the React component definition corresponding to the HTML <h1> element
type H1Def struct {
	underlying *js.Object
}

// _H1Props defines the properties for the <h1> element
type _H1Props struct {
	*BasicHTMLElement
}

func (d *H1Def) reactElement() {}

// H1 creates a new instance of a <h1> element with the provided props and
// child
func H1(props *H1Props, children ...Element) *H1Def {

	rProps := &_H1Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"h1", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &H1Def{underlying: underlying}
}
