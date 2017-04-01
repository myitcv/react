// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// NavDef is the React component definition corresponding to the HTML <nav> element
type NavDef struct {
	underlying *js.Object
}

// _NavProps defines the properties for the <nav> element
type _NavProps struct {
	*BasicHTMLElement
}

func (d *NavDef) reactElement() {}

// Nav creates a new instance of a <nav> element with the provided props and children
func Nav(props *NavProps, children ...Element) *NavDef {

	rProps := &_NavProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"nav", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &NavDef{underlying: underlying}
}
