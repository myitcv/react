// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// CodeDef is the React component definition corresponding to the HTML <code> element
type CodeDef struct {
	underlying *js.Object
}

// _CodeProps defines the properties for the <code> element
type _CodeProps struct {
	*BasicHTMLElement
}

func (d *CodeDef) reactElement() {}

// Code creates a new instance of a <code> element with the provided props
func Code(props *CodeProps, children ...Element) *CodeDef {

	rProps := &_CodeProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"code", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &CodeDef{underlying: underlying}
}
