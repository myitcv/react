// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// SpanDef is the React component definition corresponding to the HTML <p> element
type SpanDef struct {
	underlying *js.Object
}

// _SpanProps defines the properties for the <p> element
type _SpanProps struct {
	*BasicHTMLElement
}

func (d *SpanDef) reactElement() {}

// Span creates a new instance of a <p> element with the provided props and
// children
func Span(props *SpanProps, children ...Element) *SpanDef {

	rProps := &_SpanProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"span", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &SpanDef{underlying: underlying}
}
