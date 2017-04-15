// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// InputDef is the React component definition corresponding to the HTML <input> element
type InputDef struct {
	underlying *js.Object
}

// _InputProps defines the properties for the <input> element
type _InputProps struct {
	*BasicHTMLElement

	Placeholder  string `js:"placeholder"`
	Type         string `js:"type"`
	Value        string `js:"value"`
	DefaultValue string `js:"defaultValue" react:"omitempty"`
}

func (d *InputDef) reactElement() {}

// Input creates a new instance of a <input> element with the provided props
func Input(props *InputProps) *InputDef {

	rProps := &_InputProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"input", rProps}

	underlying := react.Call("createElement", args...)

	return &InputDef{
		underlying: underlying,
	}
}
