// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// LabelDef is the React component definition corresponding to the HTML <label> element
type LabelDef struct {
	underlying *js.Object
}

// _LabelProps defines the properties for the <label> element
type _LabelProps struct {
	*BasicHTMLElement

	For string `js:"htmlFor"`
}

func (d *LabelDef) reactElement() {}

// Label creates a new instance of a <label> element with the provided props and child
// element
func Label(props *LabelProps, child Element) *LabelDef {

	rProps := &_LabelProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "label", rProps, elementToReactObj(child))

	return &LabelDef{underlying: underlying}
}
