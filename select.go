// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// SelectDef is the React component definition corresponding to the HTML <select> element
type SelectDef struct {
	underlying *js.Object
}

// _SelectProps are the props for a <select> component
type _SelectProps struct {
	*BasicHTMLElement

	Value string `js:"value"`
}

func (d *SelectDef) reactElement() {}

// Select creates a new instance of a <select> element with the provided props and children
func Select(props *SelectProps, children ...*OptionDef) *SelectDef {

	rProps := &_SelectProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"select", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &SelectDef{underlying: underlying}
}
