// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// HRDef is the React component definition corresponding to the HTML <hr> element
type HRDef struct {
	underlying *js.Object
}

// _HRProps defines the properties for the <hr> element
type _HRProps struct {
	*BasicHTMLElement
}

func (d *HRDef) reactElement() {}

// HR creates a new instance of a <hr> element with the provided props
func HR(props *HRProps) *HRDef {

	rProps := &_HRProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "hr", rProps)

	return &HRDef{underlying: underlying}
}
