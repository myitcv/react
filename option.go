// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// OptionDef is the React component definition corresponding to the HTML <option> element
type OptionDef struct {
	underlying *js.Object
}

// _OptionProps defines the properties for the <option> element
type _OptionProps struct {
	*BasicHTMLElement

	Value string `js:"value"`
}

func (d *OptionDef) reactElement() {}

// Option creates a new instance of a <option> element with the provided props
func Option(props *OptionProps, child Element) *OptionDef {

	rProps := &_OptionProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"option", rProps, child}

	underlying := react.Call("createElement", args...)

	return &OptionDef{
		underlying: underlying,
	}
}
