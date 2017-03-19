// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// SpanDef is the React component definition corresponding to the HTML <p> element
type SpanDef struct {
	underlying *js.Object
}

// SpanPropsDef defines the properties for the <p> element
type SpanPropsDef struct {
	*BasicHTMLElement
}

// SpanProps creates a new instance of <p> properties, mutating the props
// by the provided initiapser
func SpanProps(f func(p *SpanPropsDef)) *SpanPropsDef {
	res := &SpanPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *SpanDef) reactElement() {}

// Span creates a new instance of a <p> element with the provided props and
// children
func Span(props *SpanPropsDef, children ...Element) *SpanDef {
	args := []interface{}{"p", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &SpanDef{underlying: underlying}
}
