// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// UlDef is the React component definition corresponding to the HTML <ul> element
type UlDef struct {
	underlying *js.Object
}

// _UlProps defines the properties for the <ul> element
type _UlProps struct {
	*BasicHTMLElement
}

func (d *UlDef) reactElement() {}

// Ul creates a new instance of a <ul> element with the provided props and <li>
// children
func Ul(props *UlProps, children ...*LiDef) *UlDef {

	rProps := &_UlProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"ul", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &UlDef{underlying: underlying}
}
