// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// NavDef is the React component definition corresponding to the HTML <nav> element
type NavDef struct {
	underlying *js.Object
}

// NavPropsDef defines the properties for the <nav> element
type NavPropsDef struct {
	*BasicHTMLElement
}

// NavProps creates a new instance of <nav> properties, mutating the props
// by the provided initialiser
func NavProps(f func(p *NavPropsDef)) *NavPropsDef {
	res := &NavPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *NavDef) reactElement() {}

// Nav creates a new instance of a <nav> element with the provided props and children
func Nav(props *NavPropsDef, children ...Element) *NavDef {
	args := []interface{}{"nav", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &NavDef{underlying: underlying}
}
