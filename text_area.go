// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// TextAreaDef is the React component definition corresponding to the HTML <textarea> element
type TextAreaDef struct {
	underlying *js.Object
}

// _TextAreaProps defines the properties for the <textarea> element
type _TextAreaProps struct {
	*BasicHTMLElement

	Rows         uint   `js:"rows"`
	Cols         uint   `js:"cols"`
	Placeholder  string `js:"placeholder"`
	Value        string `js:"value"`
	DefaultValue string `js:"defaultValue" react:"omitempty"`
}

func (d *TextAreaDef) reactElement() {}

// TextArea creates a new instance of a <textarea> element with the provided props and
// children
func TextArea(props *TextAreaProps, children ...Element) *TextAreaDef {

	rProps := &_TextAreaProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"textarea", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &TextAreaDef{underlying: underlying}
}
