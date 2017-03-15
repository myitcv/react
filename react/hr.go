// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// HRDef is the React component definition corresponding to the HTML <hr> element
type HRDef struct {
	underlying *js.Object
}

// HRPropsDef defines the properties for the <hr> element
type HRPropsDef struct {
	*BasicHTMLElement
}

// HRProps creates a new instance of <hr> properties, mutating the props
// by the provided initiahrser
func HRProps(f func(p *HRPropsDef)) *HRPropsDef {
	res := &HRPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *HRDef) reactElement() {}

// HR creates a new instance of a <hr> element with the provided props
func HR(props *HRPropsDef) *HRDef {
	underlying := react.Call("createElement", "hr", props)

	return &HRDef{underlying: underlying}
}
