// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// LabelDef is the React component definition corresponding to the HTML <label> element
type LabelDef struct {
	underlying *js.Object
}

// LabelPropsDef defines the properties for the <label> element
type LabelPropsDef struct {
	*BasicHTMLElement

	For string `js:"htmlFor"`
}

// LabelProps creates a new instance of <label> properties, mutating the props
// by the provided initialabelser
func LabelProps(f func(p *LabelPropsDef)) *LabelPropsDef {
	res := &LabelPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *LabelDef) reactElement() {}

// Label creates a new instance of a <label> element with the provided props and child
// element
func Label(props *LabelPropsDef, child Element) *LabelDef {
	underlying := react.Call("createElement", "label", props, elementToReactObj(child))

	return &LabelDef{underlying: underlying}
}
