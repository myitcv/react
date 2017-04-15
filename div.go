// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// DivDef is the React component definition corresponding to the HTML <div> element
type DivDef struct {
	underlying *js.Object
}

// _DivProps are the props for a <div> component
type _DivProps struct {
	*BasicHTMLElement
}

func (d *DivDef) reactElement() {}

// Div creates a new instance of a <div> element with the provided props and children
func Div(props *DivProps, children ...Element) *DivDef {

	rProps := &_DivProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"div", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &DivDef{underlying: underlying}
}
