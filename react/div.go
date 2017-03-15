// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// DivDef is the React component definition corresponding to the HTML <div> element
type DivDef struct {
	underlying *js.Object
}

// DivPropsDef defines the properties for the <div> element
type DivPropsDef struct {
	*BasicHTMLElement
}

// DivProps creates a new instance of <div> properties, mutating the props
// by the provided initiadivser
func DivProps(f func(p *DivPropsDef)) *DivPropsDef {
	res := &DivPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *DivDef) reactElement() {}

// Div creates a new instance of a <div> element with the provided props and children
func Div(props *DivPropsDef, children ...Element) *DivDef {
	args := []interface{}{"div", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &DivDef{underlying: underlying}
}
