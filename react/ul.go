// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// UlDef is the React component definition corresponding to the HTML <ul> element
type UlDef struct {
	underlying *js.Object
}

// UlPropsDef defines the properties for the <ul> element
type UlPropsDef struct {
	*BasicHTMLElement
}

// UlProps creates a new instance of <ul> properties, mutating the props
// by the provided initiaulser
func UlProps(f func(p *UlPropsDef)) *UlPropsDef {
	res := &UlPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *UlDef) reactElement() {}

// Ul creates a new instance of a <ul> element with the provided props and <li>
// children
func Ul(props *UlPropsDef, children ...*LiDef) *UlDef {
	args := []interface{}{"ul", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &UlDef{underlying: underlying}
}
