// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// ButtonDef is the React component definition corresponding to the HTML <button> element
type ButtonDef struct {
	underlying *js.Object
}

// _ButtonProps defines the properties for the <button> element
type _ButtonProps struct {
	*BasicHTMLElement

	Type string `js:"type"`
}

func (d *ButtonDef) reactElement() {}

// Button creates a new instance of a <button> element with the provided props
// and child
func Button(props *ButtonProps, child Element) *ButtonDef {

	rProps := &_ButtonProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"button", rProps, elementToReactObj(child)}

	underlying := react.Call("createElement", args...)

	return &ButtonDef{
		underlying: underlying,
	}
}
