// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// BRDef is the React component definition corresponding to the HTML <br> element
type BRDef struct {
	underlying *js.Object
}

// _BRProps defines the properties for the <br> element
type _BRProps struct {
	*BasicHTMLElement
}

func (d *BRDef) reactElement() {}

// BR creates a new instance of a <br> element with the provided props
func BR(props *BRProps) *BRDef {

	rProps := &_BRProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "br", rProps)

	return &BRDef{underlying: underlying}
}
