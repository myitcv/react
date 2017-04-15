// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// PreDef is the React component definition corresponding to the HTML <pre> element
type PreDef struct {
	underlying *js.Object
}

// _PreProps defines the properties for the <pre> element
type _PreProps struct {
	*BasicHTMLElement
}

func (d *PreDef) reactElement() {}

// Pre creates a new instance of a <pre> element with the provided props and
// children
func Pre(props *PreProps, children ...Element) *PreDef {

	rProps := &_PreProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"pre", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &PreDef{underlying: underlying}
}
