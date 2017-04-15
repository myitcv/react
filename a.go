// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// ADef is the React component definition corresponding to the HTML <a> element
type ADef struct {
	underlying *js.Object
}

// _APropsDef defines the properties for the <a> element
type _AProps struct {
	*BasicHTMLElement

	Target string `js:"target"`
	Href   string `js:"href"`
}

func (d *ADef) reactElement() {}

// A creates a new instance of a <a> element with the provided props and
// children
func A(props *AProps, children ...Element) *ADef {

	rProps := &_AProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"a", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &ADef{underlying: underlying}
}
