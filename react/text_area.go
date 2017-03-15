// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// TextAreaDef is the React component definition corresponding to the HTML <textarea> element
type TextAreaDef struct {
	underlying *js.Object
}

// TextAreaPropsDef defines the properties for the <textarea> element
type TextAreaPropsDef struct {
	*BasicHTMLElement

	Placeholder  string `js:"placeholder"`
	Value        string `js:"value"`
	DefaultValue string `js:"defaultValue"`
}

// TextAreaProps creates a new instance of <textarea> properties, mutating the props
// by the provided initiatextareaser
func TextAreaProps(f func(p *TextAreaPropsDef)) *TextAreaPropsDef {
	res := &TextAreaPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *TextAreaDef) reactElement() {}

// TextArea creates a new instance of a <textarea> element with the provided props and
// children
func TextArea(props *TextAreaPropsDef, children ...Element) *TextAreaDef {
	args := []interface{}{"textarea", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &TextAreaDef{underlying: underlying}
}
