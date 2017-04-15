// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

//go:generate reactGen

// PDef is the React component definition corresponding to the HTML <p> element
type PDef struct {
	underlying *js.Object
}

// _PProps are the props for a <div> component
type _PProps struct {
	*BasicHTMLElement
}

func (d *PDef) reactElement() {}

// P creates a new instance of a <p> element with the provided props and
// children
func P(props *PProps, children ...Element) *PDef {

	rProps := &_PProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"p", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &PDef{underlying: underlying}
}
