// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// LiDef is the React component definition corresponding to the HTML <li> element
type LiDef struct {
	underlying *js.Object
}

// LiPropsDef defines the properties for the <li> element
type LiPropsDef struct {
	*BasicHTMLElement

	Role string `js:"role"`
}

// LiProps creates a new instance of <li> properties, mutating the props
// by the provided initialiser
func LiProps(f func(p *LiPropsDef)) *LiPropsDef {
	res := &LiPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *LiDef) reactElement() {}

// Li creates a new instance of an <li> element with the provided props
// and children
func Li(props *LiPropsDef, children ...Element) *LiDef {
	args := []interface{}{"li", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &LiDef{underlying: underlying}
}
