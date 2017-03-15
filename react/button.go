// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// ButtonDef is the React component definition corresponding to the HTML <button> element
type ButtonDef struct {
	underlying *js.Object
}

// ButtonPropsDef defines the properties for the <button> element
type ButtonPropsDef struct {
	*BasicHTMLElement

	Type string `js:"type"`
}

// ButtonProps creates a new instance of <button> properties, mutating the props
// by the provided initiabuttonser
func ButtonProps(f func(p *ButtonPropsDef)) *ButtonPropsDef {
	res := &ButtonPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *ButtonDef) reactElement() {}

// Button creates a new instance of a <button> element with the provided props
// and child
func Button(props *ButtonPropsDef, child Element) *ButtonDef {
	args := []interface{}{"button", props, elementToReactObj(child)}

	underlying := react.Call("createElement", args...)

	return &ButtonDef{
		underlying: underlying,
	}
}
