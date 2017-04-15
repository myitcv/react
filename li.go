// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// LiDef is the React component definition corresponding to the HTML <li> element
type LiDef struct {
	underlying *js.Object
}

// _LiProps defines the properties for the <li> element
type _LiProps struct {
	*BasicHTMLElement
}

func (d *LiDef) reactElement() {}

// Li creates a new instance of an <li> element with the provided props
// and children
func Li(props *LiProps, children ...Element) *LiDef {

	rProps := &_LiProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"li", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &LiDef{underlying: underlying}
}
